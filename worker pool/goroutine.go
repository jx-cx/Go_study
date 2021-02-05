package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("worker:%d start job:%d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("worker:%d end job:%d\n", id, j)
		results <- j * 2
	}
}

func worker_poll() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	// 开启3个goroutine
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	// 5个任务
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)
	// 输出结果
	for a := 1; a <= 5; a++ {
		<-results
	}
}

// 使用goroutine和channel实现一个计算int64随机数各位数和的程序

type job struct {
	vlue int64
}
type result struct {
	job *job
	sum int64
}

var (
	jobchan    = make(chan *job, 100)
	resultchan = make(chan *result, 100)
	wg         sync.WaitGroup
)

func suijishu(sj chan<- *job) {
	// 循环生成int64的随机数，发送给jobchan
	for {
		x := rand.Int63() //创建随机数
		newjob := &job{   // 创建一个变量接收这个随机数
			vlue: x,
		}
		sj <- newjob //将随机数给sj
		time.Sleep(time.Millisecond * 500)
	}
}

func gongzuozhe(sj <-chan *job, resultChan chan<- *result) {
	// 取出随机数，计算和发送到resultChan
	for {
		job := <-sj
		sum := int64(0)
		n := job.vlue
		for n > 0 {
			sum += n % 10
			n = n / 10
		}
		newResult := &result{
			job: job,
			sum: sum,
		}
		resultChan <- newResult
	}
}

func main() {
	//worker_poll()
	wg.Add(1) // 开启一个随机数进程
	go suijishu(jobchan)
	//开启24个进程
	wg.Add(24)
	for i := 0; i < 24; i++ {
		go gongzuozhe(jobchan, resultchan)
	}
	// 从主goroutine中resultchan取出结果打印
	for result := range resultchan {
		fmt.Printf("value:%d sum:%d \n", result.job.vlue, result.sum)
	}
	wg.Wait()
}

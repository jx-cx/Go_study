package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
启动多个goroutine
在Go语言中实现并发就是这样简单，我们还可以启动多个goroutine。让我们再来一个例子： （这里使用了sync.WaitGroup来实现goroutine的同步）
*/
func f() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		r1 := rand.Int()    // 随机数 int64
		r2 := rand.Intn(10) // 0<=0 x <10
		fmt.Println(r1, r2)
	}
}

func f1(i int) {
	defer wg.Done()
	time.Sleep(time.Second * time.Duration(rand.Intn(3)))
	fmt.Println(i)
}

var wg sync.WaitGroup

func main() {

	//f()
	for i := 0; i < 10; i++ {
		wg.Add(1) // 每次执行计数器+1
		go f1(i)

	}
	wg.Wait() //等待计数器减为0

}

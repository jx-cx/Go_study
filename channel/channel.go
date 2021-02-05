package main

import (
	"fmt"
	"sync"
)

/*
单纯地将函数并发执行是没有意义的。函数与函数间需要交换数据才能体现并发执行函数的意义。

虽然可以使用共享内存进行数据交换，但是共享内存在不同的goroutine中容易发生竞态问题。为了保证数据交换的正确性，必须使用互斥量对内存进行加锁，这种做法势必造成性能问题。

Go语言的并发模型是CSP（Communicating Sequential Processes），提倡通过通信共享内存而不是通过共享内存而实现通信。

如果说goroutine是Go程序并发的执行体，channel就是它们之间的连接。channel是可以让一个goroutine发送特定值到另一个goroutine的通信机制。

Go 语言中的通道（channel）是一种特殊的类型。通道像一个传送带或者队列，总是遵循先入先出（First In First Out）的规则，保证收发数据的顺序。每一个通道都是一个具体类型的导管，也就是声明channel的时候需要为其指定元素类型。
*/

var (
	a   int
	b   chan int
	wg  sync.WaitGroup
	one sync.Once
)

//发送:ch <- 10 // 把10发送到ch中

// 接收
//x := <- ch // 从ch中接收值并赋值给变量x
//<-ch       // 从ch中接收值，忽略结果

//关闭  close(ch)

func noBufChannel() {
	fmt.Println(b) // chan 必须初始化才能使用
	//make(chan 元素类型, [缓冲大小])
	b = make(chan int) //不带缓冲区的
	//b=make(chan int,16)  // 带缓冲区的，可以放16个

	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-b
		fmt.Println("取到了B的值", x)
	}()

	b <- 10
	fmt.Println("10发送到通道b中了")
	wg.Wait()
}

func bufChannel() {
	fmt.Println(b) // chan 必须初始化才能使用
	//make(chan 元素类型, [缓冲大小])
	b = make(chan int, 1) //不带缓冲区的
	//b=make(chan int,16)  // 带缓冲区的，可以放16个

	b <- 10
	fmt.Println("10发送到通道b中了")
	//b <- 20
	//fmt.Println("10发送到通道b中了")
	x := <-b
	fmt.Println("x的值是", x)
}

func f1(ch1 chan int) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1)
}
func f2(ch1, ch2 chan int) {
	defer wg.Done()
	for {
		x, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- x * x
	}
	one.Do(func() { close(ch2) })
}

func main() {
	//noBufChannel()
	// bufChannel()
	//time.Sleep(10)
	//a := make(chan int, 100)
	//b := make(chan int, 100)
	//wg.Add(3)
	//go f1(a)
	//go f2(a, b)
	//go f2(a, b)
	//wg.Wait()
	//for ret := range b {
	//	fmt.Println(ret)
	//}

}

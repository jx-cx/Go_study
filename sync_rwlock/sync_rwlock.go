package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x      = 0
	lock   sync.Mutex
	wg     sync.WaitGroup
	rwlock sync.RWMutex
)

func reed() {
	rwlock.RLock()
	fmt.Println(x)
	time.Sleep(time.Millisecond)
	rwlock.RUnlock()
	wg.Done()
}
func write() {
	rwlock.Lock()
	x += 1
	time.Sleep(time.Millisecond * 5)
	rwlock.Unlock()
	wg.Done()
}

func main() {
	start := time.Now()
	for i := 0; i < 100; i++ {
		go write()
		wg.Add(1)
	}
	time.Sleep(time.Second)
	for i := 0; i < 1000; i++ {
		go reed()
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start))

}

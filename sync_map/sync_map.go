package main

import (
	"fmt"
	"strconv"
	"sync"
)

/*
Go语言的sync包中提供了一个开箱即用的并发安全版map–sync.Map。开箱即用表示不用像内置的map一样使用make函数初始化就能直接使用。同时sync.Map内置了诸如Store、Load、LoadOrStore、Delete、Range等操作方法。
*/

var m = make(map[string]int)
var m1 = sync.Map{}

func get(key string) int {
	return m[key]
}

func set(key string, value int) {
	m[key] = value
}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 200; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m1.Store(key, n)         // 使用Store存值（键值对
			value, _ := m1.Load(key) //使用load的方法取值
			fmt.Printf("k=:%v,v:=%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

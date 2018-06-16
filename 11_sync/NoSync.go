package main

import (
	"sync"
	"runtime"
	"fmt"
)

var (
	count int32
	wg    sync.WaitGroup
)

func main() {
	wg.Add(2)
	go incCount()
	go incCount()
	wg.Wait()
	fmt.Print(count)
}

//此函数未对全局变量count做任何保好, 所以在并发的情况下可能会出现count 为2 为3 为4的可能
func incCount() {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		value := count
		runtime.Gosched()
		value++
		count = value
	}
}

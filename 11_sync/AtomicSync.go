package main

import (
	"sync"
	"sync/atomic"
	"runtime"
	"fmt"
)

var (
	atomicCount int32
	atomicWg    sync.WaitGroup
)

func main() {
	atomicWg.Add(2)
	go atomicSyncIncCount()
	go atomicSyncIncCount()
	atomicWg.Wait()
	fmt.Print(atomicCount)
}

//
func atomicSyncIncCount() {
	defer atomicWg.Done()
	for i := 0; i < 2; i++ {
		value := atomic.LoadInt32(&atomicCount)
		runtime.Gosched()
		value++
		atomic.StoreInt32(&atomicCount, value)
	}
}

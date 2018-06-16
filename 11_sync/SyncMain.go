package main

import (
	"sync"
	"runtime"
	"fmt"
)

var (
	syncCount int32;
	syncWg    sync.WaitGroup
	mutex     sync.Mutex
)

func main() {
	syncWg.Add(2)
	go syncIncCount()
	go syncIncCount()
	syncWg.Wait()
	fmt.Print(syncCount)
}

func syncIncCount() {
	defer syncWg.Done()
	for i := 0; i < 2; i++ {
		mutex.Lock()
		value := syncCount
		runtime.Gosched()
		value++
		syncCount = value
		mutex.Unlock()
	}
}

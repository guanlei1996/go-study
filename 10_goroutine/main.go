package main

import (
	"sync"
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i < 1000; i++ {
			fmt.Println("A:", i)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 1; i < 1000; i++ {
			fmt.Println("B:", i)
		}
	}()
	wg.Wait()
}

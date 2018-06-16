## go并发资源争用

### 未进行资源保护 -> NoSync.go
```go
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

```
incCount函数未进行任何针对资源全局资源count进行任何保护处理,所以可能会出现最终结果与预期结果不符的可能.<br/>
针对这种情况, 编译时可以使用-race参数进行编译, 这样就会生成一个带有检测资源争用功能的可执行文件
```shell
go build -race
```
然后运行可执行文件
```shell
./NoSync
```
会输出检测结果以及执行结果
```shell
➜  11_sync git:(master) ✗ ./NoSync
==================
WARNING: DATA RACE
Read at 0x0000011d52cc by goroutine 7:
  main.incCount()
      /Users/lay/dev/workSpace/golang/src/github.com/guanlei1996/go-study/11_sync/NoSync.go:26 +0x6f

Previous write at 0x0000011d52cc by goroutine 6:
  main.incCount()
      /Users/lay/dev/workSpace/golang/src/github.com/guanlei1996/go-study/11_sync/NoSync.go:29 +0x8e

Goroutine 7 (running) created at:
  main.main()
      /Users/lay/dev/workSpace/golang/src/github.com/guanlei1996/go-study/11_sync/NoSync.go:17 +0x77

Goroutine 6 (finished) created at:
  main.main()
      /Users/lay/dev/workSpace/golang/src/github.com/guanlei1996/go-study/11_sync/NoSync.go:16 +0x5f
==================
4Found 1 data race(s)

```

### 对共享资源同步加锁
Go语言提供了atomic包和sync包里的一些函数对共享资源同步加锁.
#### atomic包 -> AtomicSync.go
 ```go
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


func atomicSyncIncCount() {
	defer atomicWg.Done()
	for i := 0; i < 2; i++ {
		value := atomic.LoadInt32(&atomicCount)
		runtime.Gosched()
		value++
		atomic.StoreInt32(&atomicCount, value)
	}
}

```
注意留意 atomic.LoadInt32 和 atomic.StoreInt32两个函数, 一个读取int32变量的值, 一个修改int32类型变量的值,
这两个都是原子性的操作, Go已经帮助我们在底层使用了锁的机制, 保证了共享资源的同步与安全, 所以我们可以得到正确的结果.
这时候我们再使用资源争用检测工具 go build -race检车, 也不会提示任何问题.<br/>
atomic包里还有很多原子化的函数可以保证资源在并发情况的的访问修改的问题, 比如函数
atomic.AddInt32 可以直接对一个int32类型的变量进行修改.

#### sync包 -> SyncMain.go
atomic虽然可以解决资源争用的问题, 但是所提供的函数都是比较简单的, 支持的数据类型也比较有限,所以Go语言还提供了一个sync包,
这个sync包里提供了一种互斥锁, 可以让我们自己灵活的控制哪些代码只能有一个goroutine访问.
```go
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

```
示例中, 声明了一个互斥锁 mutex, 这个互斥锁有两个方法, 一个是Lock(), 一个是Unlock().
这两个函数之间的代码就是临界区, 临界区的代码是安全的.
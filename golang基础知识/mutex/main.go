package main

import (
	"image"
	"sync"
)

/*
我们使用了一个buffered channel作为一个计数信号量， 来保证最多只有20个
goroutine会同时执行HTTP请求。 同理， 我们可以用一个容量只有1的channel来保证最多只有
一个goroutine在同一时刻访问一个共享变量。 一个只能为1和0的信号量叫做二元信号量
(binary semaphore)。
 */
var (
	sema = make(chan struct{}, 1) // a binary semaphore guarding balance
	balance int
)

func Deposit(amount int) {
	sema <- struct{}{} // acquire token
	balance = balance + amount
	<-sema // release token
}

func Balance() int {
	sema <- struct{}{} // acquire token
	b := balance
	<-sema // release token
	return b
}


//也可以使用互斥锁实现这种容量为1的channel
var (
	mu sync.Mutex // guards balance
	balance2 int
)

func Deposit2(amount int) {
	mu.Lock()
	balance = balance + amount
	mu.Unlock()
}

func Balance2() int {
	mu.Lock()
	b := balance2
	mu.Unlock()
	return b
}

// 进一步，可以使用读写所，可以并发读，写的时候互斥

var mutex sync.RWMutex
var balance3 int

func Balance3() int {
	mutex.RLock() // readers lock
	defer mutex.RUnlock()
	//mutex.Lock()，写的时候加锁方式一样
	return balance
}

// 如果初始化成本比较大的话， 那么将初始化延迟到需要的时候再去做就是一个比较好的选
//择。 如果在程序启动的时候就去做这类的初始化的话会增加程序的启动时间并且因为执行的
//时候可能也并不需要这些变量所以实际上有一些浪费。

//比如使用读写锁进行icon的初始化，但是太复杂
var mutex1 sync.RWMutex // guards icons
var icons map[string]image.Image
// Concurrency-safe.
func Icon(name string) image.Image {
	mutex1.RLock()
	if icons != nil {
		icon := icons[name]
		mutex1.RUnlock()
		return icon
	}
	mutex1.RUnlock()
	// acquire an exclusive lock
	mutex1.Lock()
	if icons == nil { // NOTE: must recheck for nil
		loadIcons()
	}
	icon := icons[name]
	mutex1.Unlock()
	return icon
}

func loadIcons()  {

}

//可以使用sync.Once来简化前面的Icon函数，每一次对Do(loadIcons)的调用都会锁定mutex， 并会检查boolean变量。 在第一次调用时， 变
//量的值是false， Do会调用loadIcons并会将boolean设置为true。 随后的调用什么都不会做，也就是首次初始化
var loadIconsOnce sync.Once
var icons2 map[string]image.Image
// Concurrency-safe.
func Icon2(name string) image.Image {
	loadIconsOnce.Do(loadIcons)
	return icons2[name]
}


/*
当我们有很多任务要同时进行时，如果并不需要关心各个任务的执行进度，那直接使用 go 关键字即可。
如果我们需要关心所有任务完成后才能往下运行时，则需要 WaitGroup 来阻塞等待这些并发任务了。
WaitGroup 如同它的字面意思，就是等待一组 goroutine 运行完成，主要有三个方法组成：

Add(delta int) ：添加任务数
Wait()：阻塞等待所有任务的完成
Done()：完成任务
 */
func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			println("hello")
		}()
	}

	wg.Wait()
}
//从上述代码可以看出，WaitGroup 的用法非常简单：使用 Add 添加需要等待的个数，使用 Done 来通知 WaitGroup 任务已完成，
//使用 Wait 来等待所有 goroutine 结束。

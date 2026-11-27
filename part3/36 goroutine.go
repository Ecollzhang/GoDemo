/*
Go 并发
并发是指程序同时执行多个任务的能力。

Go 语言支持并发，通过 goroutines 和 channels 提供了一种简洁且高效的方式来实现并发。

Goroutines：
	Go 中的并发执行单位，类似于轻量级的线程。
	Goroutine 的调度由 Go 运行时管理，用户无需手动分配线程。
	使用 go 关键字启动 Goroutine。
	Goroutine 是非阻塞的，可以高效地运行成千上万个 Goroutine。

Channel：
	Go 中用于在 Goroutine 之间通信的机制。
	支持同步和数据共享，避免了显式的锁机制。
	使用 chan 关键字创建，通过 <- 操作符发送和接收数据。

Scheduler（调度器）：
	Go 的调度器基于 GMP 模型，调度器会将 Goroutine 分配到系统线程中执行，并通过 M 和 P 的配合高效管理并发。
	G：Goroutine。
	M：系统线程（Machine）。
	P：逻辑处理器（Processor）。

Goroutine
goroutine 是轻量级线程，goroutine 的调度是由 Golang 运行时进行管理的。
goroutine 语法格式：
	go 函数名( 参数列表 )
例如：
	go f(x, y, z)
开启一个新的 goroutine:
f(x, y, z)
Go 允许使用 go 语句开启一个新的运行期线程， 即 goroutine，以一个不同的、新创建的 goroutine 来执行一个函数。 同一个程序中的所有 goroutine 共享同一个地址空间。
执行下面代码，你会看到输出的 Main 和 Hello。输出是没有固定先后顺序，因为它们是两个 goroutine 在执行
*/

/*
通道（Channel）
通道（Channel）是用于 Goroutine 之间的数据传递。
通道可用于两个 goroutine 之间通过传递一个指定类型的值来同步运行和通讯。
使用 make 函数创建一个 channel，使用 <- 操作符发送和接收数据。如果未指定方向，则为双向通道。
	ch <- v    // 把 v 发送到通道 ch
	v := <-ch  // 从 ch 接收数据
			// 并把值赋给 v
声明一个通道很简单，我们使用chan关键字即可，通道在使用前必须先创建：
ch := make(chan int)
注意：默认情况下，通道是不带缓冲区的。发送端发送数据，同时必须有接收端相应的接收数据。

以下实例通过两个 goroutine 来计算数字之和，在 goroutine 完成计算后，它会计算两个结果
*/

/*
Go 遍历通道与关闭通道
Go 通过 range 关键字来实现遍历读取到的数据，类似于与数组或切片。格式如下：
	v, ok := <-ch
如果通道接收不到数据后 ok 就为 false，这时通道就可以使用 close() 函数来关闭。
*/

/*
Select 语句
select 语句使得一个 goroutine 可以等待多个通信操作。select 会阻塞，直到其中的某个 case 可以继续执行
*/

/*
高级特性
Buffered Channel：
创建有缓冲的 Channel。
	ch := make(chan int, 2)

Context：
用于控制 Goroutine 的生命周期。
	context.WithCancel、context.WithTimeout。

Mutex 和 RWMutex：
sync.Mutex 提供互斥锁，用于保护共享资源。
	var mu sync.Mutex
	mu.Lock()
	// critical section
	mu.Unlock()
*/
package main

import (
	"fmt"
	"time"
	"sync"
)

func sayHello() {
	for i := 0; i < 5; i++ {
		fmt.Println("Hello")
		time.Sleep(100 * time.Millisecond)
	}
}

func sum(s []int, c chan int) {
    sum := 0
    for _, v := range s {
        sum += v
    }
    c <- sum // 把 sum 发送到通道 c
}


func fibonacci(n int, c chan int) {
    x, y := 0, 1
    for i := 0; i < n; i++ {
        c <- x
        x, y = y, x+y
    }
    close(c)
}

func fibonacci2(c, quit chan int) {
    x, y := 0, 1
    for {
        select {
        case c <- x:
            x, y = y, x+y
        case <-quit:
            fmt.Println("quit")
            return
        }
    }
}


func worker(id int, wg *sync.WaitGroup) {
        defer wg.Done() // Goroutine 完成时调用 Done()
        fmt.Printf("Worker %d started\n", id)
        fmt.Printf("Worker %d finished\n", id)
}


func main() {
	go sayHello() // 启动 Goroutine
	for i := 0; i < 5; i++ {
		fmt.Println("Main")
		time.Sleep(100 * time.Millisecond)
	}

	s := []int{7, 2, 8, -9, 4, 0}

    c := make(chan int)
    go sum(s[:len(s)/2], c)
    go sum(s[len(s)/2:], c)
    x, y := <-c, <-c // 从通道 c 中接收

    fmt.Println(x, y, x+y)

	c2 := make(chan int, 10)
    go fibonacci(cap(c2), c2)
    // range 函数遍历每个从通道接收到的数据，因为 c 在发送完 10 个
    // 数据之后就关闭了通道，所以这里我们 range 函数在接收到 10 个数据
    // 之后就结束了。如果上面的 c 通道不关闭，那么 range 函数就不
    // 会结束，从而在接收第 11 个数据的时候就阻塞了。
    for i2 := range c2 {
        fmt.Println(i2)
    }
	
	// 以上代码中，fibonacci goroutine 在 channel c 上发送斐波那契数列，当接收到 quit channel 的信号时退出。
	c3 := make(chan int)
    quit := make(chan int)

    go func() {
        for i := 0; i < 10; i++ {
            fmt.Println(<-c3)
        }
        quit <- 0
    }()
    fibonacci2(c3, quit)

	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1) // 增加计数器
		go worker(i, &wg)
	}
	wg.Wait() // 等待所有 Goroutine 完成
	fmt.Println("All workers done")
}


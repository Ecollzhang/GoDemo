package main

import (
	"fmt"
	"time"
)

/*
Go 语言的核心特性
Go 的设计哲学决定了它的特性集合——精简而强大。以下逐一拆解 Go 最核心的语言特性。

原生并发：Goroutine 与 Channel
并发是 Go 语言最引人注目的特性。Goroutine 是一种比操作系统线程轻量得多的执行单元——每个 Goroutine 的初始栈大小仅为约 2 KB，而操作系统线程通常在 1 MB 以上。
你可以在一台普通机器上同时运行数十万个 Goroutine，这是传统线程模型无法做到的。
Channel（通道）是 Goroutine 之间的通信管道，遵循 Go 的并发哲学：

不要通过共享内存来通信，而要通过通信来共享内存。（Don't communicate by sharing memory; share memory by communicating.）
*/

/*
Goroutine 泄漏
启动 goroutine 后，如果它永远阻塞且无法退出，就会造成 goroutine 泄漏。长期运行的服务中，goroutine 泄漏会逐渐耗尽内存。

最佳实践：
	确保每个 goroutine 都有退出路径
	使用 context.Context 传递取消信号
	使用 sync.WaitGroup 追踪 goroutine 生命周期
*/

// worker 模拟一个工作协程，从通道接收任务并处理
func worker(id int, jobs <-chan int, results chan<- int) {
	/*
	chan int：元素类型为 int 的通道
	<- 放在 chan 前面：表示这个通道只能接收（读取），不能发送（写入）
	chan int	双向	可读可写
	<-chan int	只读	只能 <-jobs 接收
	chan<- int	只写	只能 jobs <- x 发送
	*/
	// range 遍历 jobs 通道中的任务，直到通道被关闭
	for j := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, j)
		time.Sleep(time.Second) // 模拟耗时操作
		results <- j * 2 // 将结果发送到 results 通道
		fmt.Printf("Worker %d finished job %d\n", id, j)
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs) // 创建一个缓冲区大小为 numJobs 的通道，用于发送任务
	results := make(chan int, numJobs) // 创建一个缓冲区大小为 numJobs 的通道，用于接收结果

	// 启动 3 个工作协程
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// 发送 5 个任务到 jobs 通道
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // 关闭通道，通知 worker 不会再有新任务

	// 收集所有处理结果
	for a := 1; a <= numJobs; a++ {
		// fmt.Println("================================")
		fmt.Printf("Result: %d\n", <-results)
	}
}
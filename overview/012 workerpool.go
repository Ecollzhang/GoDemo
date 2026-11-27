/*
并发 Worker 池
这是 Go 并发的经典应用场景：使用限定数量的 goroutine 并行处理大批量任务。
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

// worker 从 jobs 通道接收任务，处理后将结果发送到 results 通道
// 任务处理完毕时调用 wg.Done()，通知 WaitGroup 任务已完成
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done() // 无论函数如何退出，都确保调用 Done()，即使发生 panic 也会执行
	for j := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, j)
		time.Sleep(time.Second) // 模拟耗时操作
		results <- j * 2 // 将结果发送到 results 通道
		fmt.Printf("Worker %d finished job %d\n", id, j)
	}
}

func main() {
	const numJobs = 10 // 定义总任务数
	const numWorkers = 3 // 定义工作协程数

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	var wg sync.WaitGroup // 创建 WaitGroup，用于等待所有工作协程完成

	// 启动工作协程
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1) // 每启动一个工作协程，WaitGroup 计数加 1
		go worker(w, jobs, results, &wg)
	}

	// 向通道发送所有任务
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // 关闭 jobs 通道，通知工作协程不会再有新任务

	// 等待所有工作协程完成
	go func() {
		wg.Wait()
		close(results) // 关闭 results 通道，通知主协程不会再有新结果
	}()

	// 收集所有处理结果
	for result := range results {
		fmt.Printf("Result: %d\n", result)
	}
}

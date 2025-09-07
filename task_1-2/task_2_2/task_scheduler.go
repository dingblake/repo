package main

/*题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
考察点 ：协程原理、并发任务调度。*/
import (
	"fmt"
	"sync"
	"time"
)

// Task 定义任务类型
type Task func()

// TaskResult 存储任务执行结果
type TaskResult struct {
	TaskID    int
	StartTime time.Time
	EndTime   time.Time
	Duration  time.Duration
	Error     error
}

// Scheduler 任务调度器
type Scheduler struct {
	tasks     []Task
	results   []TaskResult
	wg        sync.WaitGroup
	mutex     sync.Mutex
	startTime time.Time
}

// NewScheduler 创建新的调度器
func NewScheduler() *Scheduler {
	return &Scheduler{
		tasks:   make([]Task, 0),
		results: make([]TaskResult, 0),
	}
}

// AddTask 添加任务到调度器
func (s *Scheduler) AddTask(task Task) {
	s.tasks = append(s.tasks, task)
}

// Run 执行所有任务
func (s *Scheduler) Run() {
	s.startTime = time.Now()
	s.results = make([]TaskResult, len(s.tasks))

	for i, task := range s.tasks {
		s.wg.Add(1)
		go s.executeTask(i, task)
	}

	s.wg.Wait()
}

// executeTask 执行单个任务并记录结果
func (s *Scheduler) executeTask(id int, task Task) {
	defer s.wg.Done()

	result := TaskResult{
		TaskID:    id,
		StartTime: time.Now(),
	}

	defer func() {
		result.EndTime = time.Now()
		result.Duration = result.EndTime.Sub(result.StartTime)

		// 使用互斥锁保护共享数据
		s.mutex.Lock()
		s.results[id] = result
		s.mutex.Unlock()
	}()

	// 执行任务
	task()
}

// PrintResults 打印任务执行结果
func (s *Scheduler) PrintResults() {
	fmt.Println("任务执行结果:")
	fmt.Println("ID | 开始时间 | 结束时间 | 耗时")
	fmt.Println("----------------------------------------")

	for _, result := range s.results {
		fmt.Printf("%2d | %s | %s | %v\n",
			result.TaskID,
			result.StartTime.Format("15:04:05.000"),
			result.EndTime.Format("15:04:05.000"),
			result.Duration)
	}

	totalTime := time.Since(s.startTime)
	fmt.Printf("\n总执行时间: %v\n", totalTime)
}

func main() {
	// 创建调度器
	scheduler := NewScheduler()

	// 添加一些示例任务
	scheduler.AddTask(func() {
		time.Sleep(1 * time.Second)
		fmt.Println("任务1完成")
	})

	scheduler.AddTask(func() {
		time.Sleep(2 * time.Second)
		fmt.Println("任务2完成")
	})

	scheduler.AddTask(func() {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("任务3完成")
	})

	scheduler.AddTask(func() {
		time.Sleep(300 * time.Millisecond)
		fmt.Println("任务4完成")
	})

	// 执行所有任务
	scheduler.Run()

	// 打印结果
	scheduler.PrintResults()
}

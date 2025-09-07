package main

/*题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
考察点 ： sync.Mutex 的使用、并发数据安全。*/
import (
	"fmt"
	"sync"
)

var (
	counter int
	look    sync.Mutex
)

func main() {

	var wg sync.WaitGroup

	// 启动10个协程
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			incrementCounter()
		}()
	}

	// 等待所有协程完成
	wg.Wait()

	// 输出最终结果
	fmt.Println("最终计数器值:", counter)

}

func incrementCounter() {
	for i := 0; i < 1000; i++ {
		// 在修改共享变量前加锁
		look.Lock()
		counter++
		// 修改完成后解锁
		look.Unlock()
	}
}

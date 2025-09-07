package main

/*题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
考察点 ：原子操作、并发数据安全。*/
import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64
	var wg sync.WaitGroup

	// 启动10个协程
	for i := 0; i < 10; i++ {
		for i := 0; i < 1000; i++ {
			atomic.AddInt64(&counter, 1)
		}
	}

	// 等待所有协程完成
	wg.Wait()

	// 输出最终结果
	fmt.Println("最终计数器值:", atomic.LoadInt64(&counter))

}

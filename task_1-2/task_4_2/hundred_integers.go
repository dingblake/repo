package main

/*题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
考察点 ：通道的缓冲机制。*/
import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int, 10)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done() // 函数结束时通知 WaitGroup
		for i := 1; i <= 100; i++ {
			// 将整数发送到通道
			ch <- i
		}
		// 生产完成后关闭通道
		close(ch)
	}()
	go func() {
		defer wg.Done() // 函数结束时通知 WaitGroup
		for {
			// 从通道接收数据
			value, ok := <-ch
			if !ok {
				// 通道已关闭，退出循环
				break
			}
			// 打印接收到的整数
			fmt.Println(value)
		}
	}()
	wg.Wait()
}

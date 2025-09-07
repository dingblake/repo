package main

/*题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
考察点 ：通道的基本使用、协程间通信。*/
import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done() // 函数结束时通知 WaitGroup
		for i := 1; i <= 10; i++ {

			// 将整数发送到通道
			ch <- i

		}
		// 生产完成后关闭通道
		close(ch)
	}()
	wg.Add(1)
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

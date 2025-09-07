package main

/*题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
考察点 ： go 关键字的使用、协程的并发执行。*/
import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done() // 函数结束时通知 WaitGroup
		for i := 1; i <= 10; i = i + 2 {
			fmt.Printf("奇数: %d\n", i)

		}
	}()

	go func() {
		defer wg.Done() // 函数结束时通知 WaitGroup
		for i := 2; i <= 10; i = i + 2 {
			fmt.Printf("偶数: %d\n", i)

		}
	}()
	wg.Wait() // 等待所有 goroutine 完成
	fmt.Println("所有数字打印完成!")

}

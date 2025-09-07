package main

import (
	"fmt"
)

func main() {
	fmt.Println("--- 值增加10 ---")
	num := 5
	fmt.Printf("原始值: %d\n", num)
	// 调用函数，传递变量的地址
	Increase_ten(&num)
	// 输出修改后的值
	fmt.Printf("修改后的值: %d\n", num)
	fmt.Println("-------------------")
	//----------------------------------------------------------
	fmt.Println("--- 每个元素乘以2 ---")
	nums := []int{5, 1, 3, 4, 6}
	fmt.Printf("原始值: %d\n", nums)
	// 调用函数，传递变量的地址
	upnum := MultiplySliceByTwo(nums)
	// 输出修改后的值
	fmt.Printf("修改后的值: %d\n", upnum)
	nums1 := []int{11, 12, 13, 4, 6}
	fmt.Printf("原始值: %d\n", nums1)
	doubleSliceElements(&nums1)
	fmt.Printf("修改后的值: %d\n", nums1)
	fmt.Println("-------------------")
	//----------------------------------------------------------

}

/*
题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
考察点 ：指针的使用、值传递与引用传递的区别。
*/
func Increase_ten(num *int) {
	*num += 10 // 解引用指针并增加10
}

/*
题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
考察点 ：指针运算、切片操作。
*/
func MultiplySliceByTwo(nums []int) []int {
	result := []int{}
	for _, char := range nums {
		result = append(result, char*2)
	}
	return result
}

func doubleSliceElements(slicePtr *[]int) {
	// 解引用指针获取切片
	slice := *slicePtr

	// 遍历切片，将每个元素乘以2
	for i := 0; i < len(slice); i++ {
		slice[i] *= 2
	}
}

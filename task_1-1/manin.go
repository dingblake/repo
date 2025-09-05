package main

import (
	"fmt"
	"sort"
)

func main() {
	num := []int{1, 2, 2, 3, 3, 4, 4, 5, 5, 6}
	num1 := []int{1, 3, 3, 4, 33, 32, 32}
	fmt.Println("--- 不重复的元素 ---")

	fmt.Print("输出值为", findSingleNumbers(num))

	fmt.Print("输出值为", findSingleNumbers(num1))
	fmt.Println("-------------------")
	//----------------------------------------------------------
	fmt.Println("--- 回文数判断 ---")
	a := 12345654321
	b := 12034
	fmt.Print("输入值为", a)
	fmt.Print("输出值为", palindromic_number(a))
	fmt.Print("输入值为", b)
	fmt.Print("输出值为", palindromic_number(b))
	fmt.Println("-------------------")
	//----------------------------------------------------------
	fmt.Println("--- 括号判断 ---")
	bracket := "(){}[]"
	bracket1 := "({]"
	fmt.Print("输入值为", bracket)
	fmt.Print("输出值为", isValid(bracket))
	fmt.Print("输入值为", bracket1)
	fmt.Print("输出值为", isValid(bracket1))
	fmt.Println("-------------------")
	//----------------------------------------------------------
	fmt.Println("--- 最长公共前缀 ---")
	s1 := []string{"aa", "a1", "a3", "a2", "abc"}
	s2 := []string{}
	fmt.Print("输入值为", s1)
	fmt.Print("输出值为", max_common_Prefix(s1))
	fmt.Print("输入值为", s2)
	fmt.Print("输出值为", max_common_Prefix(s2))
	fmt.Println("-------------------")
	// 	//----------------------------------------------------------
	fmt.Println("--- 加一  ---")
	digits := []int{1, 2, 3}
	digits1 := []int{9, 9, 9}
	fmt.Print("输入值为", digits)
	fmt.Print("输出值为", plus_one(digits))
	fmt.Print("输入值为", digits1)
	fmt.Print("输出值为", plus_one(digits1))
	fmt.Println("-------------------")
	// 	//----------------------------------------------------------
	fmt.Println("--- 删除重复 ---")
	n1 := []int{1, 2, 2, 3, 4, 4, 5, 6, 5, 8, 7, 7, 9, 6}
	n2 := []int{1, 2, 2, 3, 4, 4, 5, 6, 5, 8, 7, 7, 9, 6}
	fmt.Print("输入值为", n1)
	fmt.Print("输出值为", remove_duplicates(n1))
	fmt.Print("输入值为", n2)
	fmt.Print("输出值为", remove_duplicates_2(n2))
	fmt.Println("-------------------")
	//----------------------------------------------------------
	fmt.Println("--- 合并区间 ---")
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	intervals1 := [][]int{{4, 7}, {1, 5}, {8, 10}, {0, 8}}
	fmt.Print("输入值为", intervals)
	fmt.Print("输出值为", merge(intervals))
	fmt.Print("输入值为", intervals1)
	fmt.Print("输出值为", merge(intervals1))
	fmt.Println("-------------------")
	// ----------------------------------------------------------
	fmt.Println("--- 两数之和 ---")
	nums3 := []int{1, 2, 5, 6, 8, 9, 6}
	target1 := 12
	fmt.Printf("输入数组: %v\n", nums3)
	fmt.Printf("输入target: %v\n", target1)
	result := Sum_of_Two_Numbers(nums3, target1)
	fmt.Printf("输出: %v\n", result)
	result1 := Sum_of_Two_Numbers1(nums3, target1)
	fmt.Printf("输出: %v\n", result1)
	fmt.Println("-------------------")
	// ----------------------------------------------------------
}

/*只出现一次的数字：给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。
找出那个只出现了一次的元素。可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，
例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。*/

func findSingleNumbers(nums []int) []int {
	// 使用map记录每个数字出现的次数
	countMap := make(map[int]int)

	// 遍历数组，统计每个数字的出现次数
	for _, num := range nums {
		countMap[num]++
	}

	// 找出所有只出现一次的数字
	result := []int{}
	for num, count := range countMap {
		if count == 1 {
			result = append(result, num)
		}
	}

	return result
}

/* 给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
例如，121 是回文，而 123 不是。*/

func palindromic_number(num int) bool {
	reversed := 0
	if num <= 0 || (num%10 == 0 && num != 0) {
		return false
	}
	for num > reversed {
		reversed = reversed*10 + num%10
		num /= 10
	}
	return num == reversed || num == reversed/10

}

/* 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
   有效字符串需满足：
   左括号必须用相同类型的右括号闭合。
   左括号必须以正确的顺序闭合。
   每个右括号都有一个对应的相同类型的左括号。*/

func isValid(s string) bool {
	stack := []rune{}
	for _, char := range s {
		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if (char == ')' && top != '(') || (char == '}' && top != '{') || (char == ']' && top != '[') {
				return false
			}
		}
	}
	return len(stack) == 0
}

/*
编写一个函数来查找字符串数组中的最长公共前缀。
如果不存在公共前缀，返回空字符串 ""。
*/
func max_common_Prefix(s1 []string) string {

	if len(s1) == 0 {
		return ""
	}
	result := s1[0]
	for i := 1; i < len(s1); i++ {
		for len(result) > len(s1[i]) || result != s1[i][:len(result)] {
			result = result[:len(result)-1]
			if result == "" {
				return ""
			}
		}
	}

	return result
}

/*
给定一个表示 大整数 的整数数组 digits，其中 digits[i] 是整数的第 i 位数字。
这些数字按从左到右，从最高位到最低位排列。这个大整数不包含任何前导 0。
将大整数加 1，并返回结果的数字数组。
*/
func plus_one(nums []int) []int {
	one := 1
	n := len(nums)
	for i := n - 1; i >= 0; i-- {
		sum := nums[i] + one
		if sum == 10 {
			nums[i] = 0
			one = 1
		} else {
			nums[i] = sum
			one = 0
		}
	}
	if one == 1 {
		nums = append([]int{1}, nums...)
	}

	return nums
}

/*给你一个 非严格递增排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致
。然后返回 nums 中唯一元素的个数。
给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
可以使用双指针法，一个慢指针 i 用于记录不重复元素的位置，一个快指针 j 用于遍历数组，当 nums[i] 与 nums[j] 不相等时，将 nums[j] 赋值给 nums[i + 1]，并将 i 后移一位。
*/
//无顺序排序
func remove_duplicates(nums []int) []int {
	array := []int{}
	countMap := make(map[int]int)
	for _, num := range nums {
		countMap[num]++
	}
	for num, count := range countMap {
		if count > 1 {
			array = append(array, num)
		} else {
			array = append(array, num)
		}
	}
	return array
}

// 修改过有顺序排序
func remove_duplicates_2(nums []int) []int {
	seen := make(map[int]bool)
	result := []int{}
	for _, num := range nums {
		if !seen[num] {
			seen[num] = true
			result = append(result, num)
		}
	}
	return result
}

/*
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
可以先对区间数组按照区间的起始位置进行排序，然后使用一个切片来存储合并后的区间，遍历排序后的区间数组，
将当前区间与切片中最后一个区间进行比较，如果有重叠，则合并区间；如果没有重叠，则将当前区间添加到切片中。
*/
func merge(intervals [][]int) [][]int {
	// 如果区间数组为空，直接返回
	if len(intervals) == 0 {
		return intervals
	}

	// 根据区间的起始点进行排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 初始化结果数组，包含第一个区间
	result := [][]int{intervals[0]}

	// 从第二个区间开始遍历
	for i := 1; i < len(intervals); i++ {
		last := result[len(result)-1] // 结果数组的最后一个区间
		current := intervals[i]       // 当前区间

		// 检查当前区间是否与最后一个区间重叠
		if current[0] <= last[1] {
			// 重叠时，合并区间：更新最后一个区间的结束点
			if current[1] > last[1] {
				last[1] = current[1]
			}
		} else {
			// 不重叠时，将当前区间加入结果数组
			result = append(result, current)
		}
	}

	return result
}

/*
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
你可以假设每种输入只会对应一个答案，并且你不能使用两次相同的元素。
你可以按任意顺序返回答案。
*/
//自己的思路
func Sum_of_Two_Numbers(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{nums[i], nums[j]}
			}
		}
	}
	return nil
}

// 询问查询ai的思路
func Sum_of_Two_Numbers1(nums []int, target int) []int {
	// 创建一个 map，用于存储元素值到索引的映射
	numMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		// 检查 complement 是否存在于 map 中
		if index, exists := numMap[complement]; exists {
			// 如果存在，返回 complement 的索引和当前索引
			return []int{nums[index], nums[i]}
		}
		// 将当前元素存入 map
		numMap[num] = i
	}
	// 根据假设，总会有一个解，所以这里不会执行到
	return nil
}

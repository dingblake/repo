package main

import "fmt"

/*题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
考察点 ：组合的使用、方法接收者。*/
//import  "fmt"

type Person struct {
	// 包含 Name 和 Age 字段
	name string
	age  int
}
type Employee struct {
	Person     // 匿名嵌入 Person 结构体
	EmployeeID string
}

func (e Employee) PrintInfo() {
	// 输出员工信息，包括姓名、年龄和员工ID
	// 可以直接访问 e.Name 和 e.Age（通过组合）
	// 以及 e.EmployeeID
	fmt.Printf("姓名: %s\n", e.name) // 可以直接访问嵌入的 Person 字段
	fmt.Printf("年龄: %d\n", e.age)
	fmt.Printf("员工ID: %s\n", e.EmployeeID)

}

func main() {
	emp := Employee{
		Person: Person{
			// 初始化 Name 和 Age
			name: "张三",
			age:  30,
		},
		EmployeeID: "001",
	}

	// 调用 PrintInfo() 方法
	emp.PrintInfo()
}

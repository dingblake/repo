package main

/*题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
考察点 ：接口的定义与实现、面向对象编程风格。*/
import (
	"fmt"
	"math"
)

// Shape 接口定义
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle (矩形)结构体定义
type Rectangle struct {
	Width, Height float64
}

// Rectangle 实现 Shape 接口的 Area 方法
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Rectangle 实现 Shape 接口的 Perimeter 方法
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Circle 结构体定义
type Circle struct {
	Radius float64
}

// Circle 实现 Shape 接口的 Area 方法
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Circle 实现 Shape 接口的 Perimeter 方法
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// 辅助函数：打印形状的信息
func printShapeInfo(s Shape) {
	fmt.Printf("面积: %.2f, 周长: %.2f\n", s.Area(), s.Perimeter())
}

func main() {
	// 创建 Rectangle 实例
	rect := Rectangle{Width: 10, Height: 5}
	fmt.Println("矩形:")
	fmt.Printf("宽度: %.2f, 高度: %.2f\n", rect.Width, rect.Height)
	printShapeInfo(rect)
	fmt.Println()

	// 创建 Circle 实例
	circle := Circle{Radius: 4}
	fmt.Println("圆形:")
	fmt.Printf("半径: %.2f\n", circle.Radius)
	printShapeInfo(circle)
	fmt.Println()

	// 使用 Shape 接口类型的切片
	fmt.Println("使用 Shape 接口切片:")
	shapes := []Shape{
		Rectangle{Width: 2, Height: 3},
		Circle{Radius: 2.5},
		Rectangle{Width: 4, Height: 4},
		Circle{Radius: 1},
	}

	for i, shape := range shapes {
		fmt.Printf("形状 %d: ", i+1)
		printShapeInfo(shape)
	}
}

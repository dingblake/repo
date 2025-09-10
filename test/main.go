package main

import (
	"fmt"
)

func main() {
	// var s string = "Hello, world!"
	// var bytes []byte = []byte(s)
	// fmt.Println("convert \"Hello, world!\" to bytes: ", bytes)
	// fmt.Println(r1)
	// fmt.Println(r2)
	// fmt.Println(s)
	// fmt.Println(runes)
	// 	var s1 string = "Hello\nworld!\n"
	// 	var s2 string = `Hello
	// world!
	// `
	// 	fmt.Println(s1 == s2)
	// var s string = "Go语言"
	// var bytes []byte = []byte(s)
	// var runes []rune = []rune(s)

	// fmt.Println("string length: ", len(s))
	// fmt.Println("bytes length: ", len(bytes))
	// fmt.Println("runes length: ", len(runes))
	// fmt.Println("string sub: ", s[0:8])
	// fmt.Println("bytes sub: ", string(bytes[0:7]))
	// fmt.Println("runes sub: ", string(runes[0:2]))
	// var s1 string = "Hello"
	// var zero int
	// var b1 = true

	// var (
	// 	i  int = 123
	// 	b2 bool
	// 	s2 = "test"
	// )

	// var (
	// 	group = 2
	// )
	// fmt.Println(s1)
	// fmt.Println(zero)
	// fmt.Println(b1)
	// fmt.Println(i)
	// fmt.Println(b2)
	// fmt.Println(s2)
	// fmt.Println(group)
	// a1, b3 := fun1()
	// fmt.Println(a1, b3)
	// method()
	// a1, a2 := fun1()
	// fmt.Println(a1, a2)\
	// var zz *string
	// var p1 *string
	// var p2 *int
	// i := 1
	// s := "hello"
	// p2 = &i
	// p1 = &s
	// p3 := &p2
	// fmt.Println(p3)
	// fmt.Println(p1)
	// fmt.Println(p2)
	// fmt.Println(zz)
	// var p1 *int
	// i := 1
	// p1 = &i
	// fmt.Println(*p1 == i)
	// *p1 = 2
	// fmt.Println(i)

	// 	a := 2
	// 	var p *int
	// 	fmt.Println(&a)
	// 	p = &a
	// 	fmt.Println(p, &a)

	// 	var pp **int
	// 	pp = &p
	// 	fmt.Println(pp, p)
	// 	**pp = 3
	// 	fmt.Println(pp, *pp, p)
	// 	fmt.Println(**pp, *p)
	// 	fmt.Println(a, &a)
	// 	type Person struct {
	// 		Name  string
	// 		Age   int
	// 		Call  func() byte
	// 		Map   map[string]string
	// 		Ch    chan string
	// 		Arr   [32]uint8
	// 		Slice []interface{}
	// 		Ptr   *int
	// 		once  sync.Once
	// 	}

	// }

	// type A struct {
	// 	a  string
	// 	c1 int
	// }

	// type B struct {
	// 	A
	// 	b string
	// }
	// type c1 struct {
	// 	c12 int
	// }

	// type C struct {
	// 	A
	// 	B
	// 	c1

	// 	a string
	// 	b string
	// 	c string
	// }

	// func fun1() (a int, b string) {
	// 	return 1, "2"
	// }

	// func method1() {
	// 	// 方式1，类型推导，用得最多
	// 	a := 1
	// 	a1 := 3
	// 	// 方式2，完整的变量声明写法
	// 	var b int = 2
	// 	// 方式3，仅声明变量，但是不赋值，
	// 	var c int
	// 	fmt.Println(a, b, c, a1)
	// }

	// // 方式4，直接在返回值中声明
	// func method2() (a int, b string) {
	// 	// 这种方式必须声明return关键字
	// 	// 并且同样不需要使用，并且也不用必须给这种变量赋值
	// 	return 1, "test"
	// }

	// func method3() (a int, b string) {
	// 	a = 1
	// 	b = "test"
	// 	return
	// }

	// func method4() (a int, b string) {
	// 	return
	// }

	// var a, b, c int = 1, 2, 3

	// var e, f, g int

	// var h, i, j = 1, 2, "test"

	// func method() {
	// 	var k, l, m int = 1, 2, 3
	// 	var n, o, p int
	// 	q, r, s := 1, 2, "test"
	// 	fmt.Println(k, l, m, n, o, p, q, r, s)
	// }

	// type A struct {
	// 	a string
	// }

	// func (a A) string() string {
	// 	return a.a
	// }

	// func (a A) stringA() string {
	// 	return a.a
	// }

	// func (a A) setA(v string) {
	// 	a.a = v
	// }

	// func (a *A) stringPA() string {
	// 	return a.a
	// }

	// func (a *A) setPA(v string) {
	// 	a.a = v
	// }

	// type B struct {
	// 	A
	// 	b string
	// }

	// func (b B) string() string {
	// 	return b.b
	// }

	// func (b B) stringB() string {
	// 	return b.b
	// }

	// type C struct {
	// 	B
	// 	a string
	// 	b string
	// 	c string
	// 	d []byte
	// }

	// func (c C) string() string {
	// 	return c.c
	// }

	// func (c C) modityD() {
	// 	c.d[2] = 3
	// }

	// func callStructMethod() {
	// 	var a A
	// 	a = A{
	// 		a: "a",
	// 	}
	// 	a.string()
	// }

	// func NewC() C {
	// 	return C{
	// 		B: B{
	// 			A: A{
	// 				a: "ba",
	// 			},
	// 			b: "b",
	// 		},
	// 		a: "ca",
	// 		b: "cb",
	// 		c: "c",
	// 		d: []byte{1, 2, 3},
	// 	}
	// }

	// func main() {
	// 	c := NewC()
	// 	cp := &c
	// 	fmt.Println(c.string())
	// 	fmt.Println(c.stringA())
	// 	fmt.Println(c.stringB())

	// 	fmt.Println(cp.string())
	// 	fmt.Println(cp.stringA())
	// 	fmt.Println(cp.stringB())

	// 	c.setA("1a")
	// 	fmt.Println("------------------c.setA")
	// 	fmt.Println(c.A.a)
	// 	fmt.Println(cp.A.a)

	// 	cp.setA("2a")
	// 	fmt.Println("------------------cp.setA")
	// 	fmt.Println(c.A.a)
	// 	fmt.Println(cp.A.a)

	// 	c.setPA("3a")
	// 	fmt.Println("------------------c.setPA")
	// 	fmt.Println(c.A.a)
	// 	fmt.Println(cp.A.a)

	// 	cp.setPA("4a")
	// 	fmt.Println("------------------cp.setPA")
	// 	fmt.Println(c.A.a)
	// 	fmt.Println(cp.A.a)

	// 	cp.modityD()
	// 	fmt.Println("------------------cp.modityD")
	// 	fmt.Println(cp.d)
	// 	const  a int =123

	// }
	// type Gender string

	// const (
	// 	Male   Gender = "Male"
	// 	Female Gender = "Female"
	// )

	// func (g *Gender) String() string {
	// 	switch *g {
	// 	case Male:
	// 		return "Male"
	// 	case Female:
	// 		return "Female"
	// 	default:
	// 		return "Unknown"
	// 	}
	// }

	// func (g *Gender) IsMale() bool {
	// 	return *g == Male
	// }

	// type ConnState int

	// const (
	// 	StateNew ConnState = iota
	// 	StateActive
	// 	StateIdle
	// 	StateHijacked
	// 	StateClosed
	// )

	// // src/time/time.go
	// type Month int

	// const (
	// 	January Month = 1 + iota
	// 	February
	// 	March
	// 	April
	// 	May
	// 	June
	// 	July
	// 	August
	// 	September
	// 	October
	// 	November
	// 	December
	// )

	// func main() {
	// 	// fmt.Println(January)
	// 	// fmt.Println(February)
	// 	// fmt.Println(March)
	// 	// fmt.Println(April)
	// 	// fmt.Println(May)
	// 	// fmt.Println(June)
	// 	// fmt.Println(July)

	// 	// a, b := 2, 10
	// 	// sum := a + b
	// 	// sub := a - b
	// 	// mul := a * b
	// 	// div := a / b
	// 	// mod := a % b

	// 	// fmt.Println(sum, sub, mul, div, mod)
	// 	// fmt.Println(a == b)
	// 	// fmt.Println(a != b)
	// 	// fmt.Println(a > b)
	// 	// fmt.Println(a < b)
	// 	// fmt.Println(a >= b)
	// 	// fmt.Println(a <= b)

	// 	a1 := true
	// 	b2 := false
	// 	fmt.Println(a1 && b2)
	// 	fmt.Println(a1 || b2)
	// 	fmt.Println(!(a1 && b2))

	// 	fmt.Println(0 & 0)
	// 	fmt.Println(0 | 0)
	// 	fmt.Println(0 ^ 0)

	// 	fmt.Println(0 & 1)
	// 	fmt.Println(0 | 1)
	// 	fmt.Println(0 ^ 1)

	// 	fmt.Println(1 & 1)
	// 	fmt.Println(1 | 1)
	// 	fmt.Println(1 ^ 1)

	// 	fmt.Println(1 & 0)
	// 	fmt.Println(1 | 0)
	// 	fmt.Println(1 ^ 0)

	// }
	// func main() {
	// 	a, b := 1, 2
	// 	var c int
	// 	c = a + b
	// 	fmt.Println("c = a + b, c =", c)

	// 	plusAssignment(c, a)
	// 	subAssignment(c, a)
	// 	mulAssignment(c, a)
	// 	divAssignment(c, a)
	// 	modAssignment(c, a)
	// 	leftMoveAssignment(c, a)
	// 	rightMoveAssignment(c, a)
	// 	andAssignment(c, a)
	// 	orAssignment(c, a)
	// 	norAssignment(c, a)
	// }

	// func plusAssignment(c, a int) {
	// 	c += a // c = c + a
	// 	fmt.Println("c += a, c =", c)
	// }

	// func subAssignment(c, a int) {
	// 	c -= a // c = c - a
	// 	fmt.Println("c -= a, c =", c)
	// }

	// func mulAssignment(c, a int) {
	// 	c *= a // c = c * a
	// 	fmt.Println("c *= a, c =", c)
	// }

	// func divAssignment(c, a int) {
	// 	c /= a // c = c / a
	// 	fmt.Println("c /= a, c =", c)
	// }

	// func modAssignment(c, a int) {
	// 	c %= a // c = c % a
	// 	fmt.Println("c %= a, c =", c)
	// }

	// func leftMoveAssignment(c, a int) {
	// 	c <<= a // c = c << a
	// 	fmt.Println("c <<= a, c =", c)
	// }

	// func rightMoveAssignment(c, a int) {
	// 	c >>= a // c = c >> a
	// 	fmt.Println("c >>= a, c =", c)
	// }

	// func andAssignment(c, a int) {
	// 	c &= a // c = c & a
	// 	fmt.Println("c &= a, c =", c)
	// }

	// func orAssignment(c, a int) {
	// 	c |= a // c = c | a
	// 	fmt.Println("c |= a, c =", c)
	// }

	// func norAssignment(c, a int) {
	// 	c ^= a // c = c ^ a
	// 	fmt.Println("c ^= a, c =", c)
	// }

	// func main() {
	// 	a := 4
	// 	var ptr *int
	// 	fmt.Println(a)

	//		ptr = &a
	//		fmt.Printf("*ptr 为 %d\n", *ptr)
	//	}
	// func main() {
	// 	var a int = 21
	// 	var b int = 10
	// 	var c int = 16
	// 	var d int = 5
	// 	var e int

	// 	e = (a + b) * c / d // ( 31 * 16 ) / 5
	// 	fmt.Printf("(a + b) * c / d 的值为 : %d\n", e)

	// 	e = ((a + b) * c) / d // ( 31 * 16 ) / 5
	// 	fmt.Printf("((a + b) * c) / d 的值为  : %d\n", e)

	// 	e = (a + b) * (c / d) // 31 * (16/5)
	// 	fmt.Printf("(a + b) * (c / d) 的值为  : %d\n", e)

	// 	// 21 + (160/5)
	// 	e = a + (b*c)/d
	// 	fmt.Printf("a + (b * c) / d 的值为  : %d\n", e)

	//		// 2 & 2 = 2; 2 * 3 = 6; 6 << 1 = 12; 3 + 4 = 7; 7 ^ 3 = 4;4 | 12 = 12
	//		f := 3 + 4 ^ 3 | 2&2*3<<1
	//		fmt.Println(f == 12)
	//		fmt.Println(f)
	//	}
	//
	//	func main() {
	//		var a int = 10
	//		if b := 1; a > 10 {
	//			b = 2
	//			// c = 2
	//			fmt.Println("a > 10")
	//		} else if c := 3; b > 1 {
	//			b = 3
	//			fmt.Println("b > 1")
	//		} else {
	//			fmt.Println("其他")
	//			if c == 3 {
	//				fmt.Println("c == 3")
	//			}
	//			fmt.Println(b)
	//			fmt.Println(c)
	//		}

	// a := "test string"
	// b := "s"

	// // 1. 基本用法
	// switch b {
	// case "test":
	// 	fmt.Println("a = ", a)
	// case "s":
	// 	fmt.Println("b = ", b)
	// case "t", "test string": // 可以匹配多个值，只要一个满足条件即可
	// 	fmt.Println("catch in a test, a = ", a)
	// case "n":
	// 	fmt.Println("a = not")
	// default:
	// 	fmt.Println("default case")
	// }

	// switch a {
	// case "test":
	// 	fmt.Println("a = ", a)
	// case "s":
	// 	fmt.Println("a = ", a)
	// case "t", "test string": // 可以匹配多个值，只要一个满足条件即可
	// 	fmt.Println("catch in a test, a = ", a)
	// case "n":
	// 	fmt.Println("a = not")
	// default:
	// 	fmt.Println("default case")
	// }

	// 变量b仅在当前switch代码块内有效
	// switch b := 5; b {
	// case 1:
	// 	fmt.Println("b = 1")
	// case 2:
	// 	fmt.Println("b = 2")
	// case 3, 4:
	// 	fmt.Println("b = 3 or 4")
	// case 5:
	// 	fmt.Println("b = 5")
	// default:
	// 	fmt.Println("b = ", b)
	// }
	// 不指定判断变量，直接在case中添加判定条件
	// b := 5
	// a := "s"
	// switch {
	// case a == "t":
	// 	fmt.Println("a = t")
	// case b == 3:
	// 	fmt.Println("b = 5")
	// case b == 5 && a == "test string":
	// 	fmt.Println("a = test string; or b = 5")
	// default:
	// 	fmt.Println("default case")
	// }
	// type MyCustomType int
	// var d interface{}
	// // var e byte = 1
	// d = MyCustomType(1)
	// //CustomType = 0
	// switch t := d.(type) {
	// case byte:
	// 	fmt.Println("d is byte type, ", t)
	// case *byte:
	// 	fmt.Println("d is byte point type, ", t)
	// case *int:
	// 	fmt.Println("d is int type, ", t)
	// case *string:
	// 	fmt.Println("d is string type, ", t)
	// case *MyCustomType:
	// 	fmt.Println("d is CustomType pointer type, ", t)
	// case MyCustomType:
	// 	fmt.Println("d is CustomType type, ", t)
	// default:
	// 	fmt.Println("d is unknown type, ", t)
	// }

	//type MyCustomType int

	// var d interface{}
	// d = MyCustomType(1) // 将 d 设置为 MyCustomType 类型

	// switch t := d.(type) {
	// case byte:
	// 	fmt.Println("d is byte type, ", t)
	// case *byte:
	// 	fmt.Println("d is byte point type, ", t)
	// case int:
	// 	fmt.Println("d is int type, ", t)
	// case *int:
	// 	fmt.Println("d is int pointer type, ", t)
	// case string:
	// 	fmt.Println("d is string type, ", t)
	// case *string:
	// 	fmt.Println("d is string pointer type, ", t)
	// case MyCustomType: // 使用类型名称而不是变量名
	// 	fmt.Println("d is MyCustomType type, ", t)
	// case *MyCustomType: // 使用类型名称而不是变量名
	// 	fmt.Println("d is MyCustomType pointer type, ", t)
	// default:
	// 	fmt.Println("d is unknown type, ", t)
	// }

	// for i := 0; i < 10; i++ {
	// 	fmt.Println("方式1，第", i+1, "次循环")
	// }

	//r a int = 10
	//		if b := 1; a > 10 {
	//			b = 2
	//			// c = 2
	//			fmt.Println("a > 10")
	//		} else if c := 3; b > 1 {
	//			b = 3
	//			fmt.Println("b > 1")
	//		} else {
	//			fmt.Println("其他")
	//			if c == 3 {
	//				fmt.Println("c == 3")
	//			}
	//			fmt.Println(b)
	//			fmt.Println(c)
	//		}

	// // 方式2
	// b := 1
	// for b < 10 {
	// 	fmt.Println("方式2，第", b, "次循环")
	// 	if b < 9 {
	// 		b = b + 1
	// 	} else {
	// 		b = b - 1
	// 	}

	// }

	// 方式3，无限循环
	// ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Second*2))
	// fmt.Println(ctx)
	// var started bool
	// var stopped atomic.Bool
	// for {
	// 	if !started {
	// 		started = true
	// 		go func() {
	// 			for {
	// 				select {
	// 				case <-ctx.Done():
	// 					fmt.Println("ctx done")
	// 					stopped.Store(true)
	// 					return
	// 				}
	// 			}
	// 		}()
	// 	}
	// 	fmt.Println("main")
	// 	if stopped.Load() {
	// 		break
	// 	}
	// }
	// 遍历数组
	// var a [10]string
	// a[0] = "Hello"
	// fmt.Println(a)
	// for i := range a {
	// 	fmt.Println("当前下标：", i)
	// }
	// for i, e := range a {
	// 	fmt.Println("a[", i, "] = ", e)
	// }

	// 遍历切片
	// s := make([]string, 10)

	// s[0] = "Hello"
	// fmt.Println(s)
	// for i := range s {
	// 	fmt.Println("当前下标：", i)
	// }
	// for i, e := range s {
	// 	fmt.Println("s[", i, "] = ", e)
	// }

	// m := make(map[string]string)
	// m["a"] = "Hello, 1"
	// m["b"] = "Hello, 2"
	// m["c"] = "Hello, 3"
	// fmt.Println(m)
	// for i := range m {
	// 	fmt.Println("当前key：", i)
	// 	for k, v := range i {
	// 		fmt.Println("m[", k, "] = ", v)
	// 	}
	// }
	// for k, v := range m {
	// 	fmt.Println("m[", k, "] = ", v)
	// }

	// for i := 1; i < 5; i++ {
	// 	if i == 3 {
	// 		break
	// 	}
	// 	fmt.Println("第", i, "次循环")
	// }

	//     // 中断switch
	// switch i := 2; i {
	// case 1:
	// 	fmt.Println("进入case 1")
	// 	if i == 1 {
	// 		break
	// 	}
	// 	fmt.Println("i等于1")
	// case 2:
	// 	fmt.Println("i等于2")
	// default:
	// 	fmt.Println("default case")
	// }

	//     // for {
	// 中断select
	// select {
	// case <-time.After(time.Second * 2):
	// 	fmt.Println("过了2秒")
	// case <-time.After(time.Second):
	// 	fmt.Println("进过了1秒")
	// 	if true {
	// 		break
	// 	}
	// 	fmt.Println("break 之后")
	// }
	// }

	//     // 不使用标记
	// 	for i := 1; i <= 3; i++ {
	// 		fmt.Printf("不使用标记,外部循环, i = %d\n", i)
	// 		for j := 5; j <= 10; j++ {
	// 			fmt.Printf("不使用标记,内部循环 j = %d\n", j)
	// 			break
	// 		}
	// 	}

	// 	// 使用标记
	// outter:
	// 	for i := 1; i <= 3; i++ {
	// 		fmt.Printf("使用标记,外部循环, i = %d\n", i)
	// 		for j := 5; j <= 10; j++ {
	// 			fmt.Printf("使用标记,内部循环 j = %d\n", j)
	// 			break outter
	// 		}
	// 	}
	// 	gotoPreset := false
	// preset:
	// 	a := 5
	// process:
	// 	//fmt.Println(gotoPreset, a)
	// 	if a > 0 {
	// 		a--
	// 		fmt.Println("当前a的值为：", a)
	// 		goto process
	// 	} else if a <= 0 {
	// 		// elseProcess:
	// 		goto post
	// 		if !gotoPreset {
	// 			gotoPreset = true
	// 			goto preset
	// 		} else {
	// 			goto post
	// 		}
	// 	}

	// post:
	// 	fmt.Println("main将结束，当前a的值为：", a)
	// 	custom()
	// }

	// type A struct {
	// 	i int
	// }

	// //	func custom() {
	// //		fmt.Println("Hello, world!")
	// //	}
	// func (a *A) add(v int) int {
	// 	a.i += v
	// 	return a.i
	// }

	// // 声明函数变量
	// var function1 func(int) int

	// // 声明闭包
	// var squart2 func(int) int = func(p int) int {
	// 	p *= p
	// 	return p
	// }

	// func main() {
	// 	a := A{1}
	// 	// 把方法赋值给函数变量
	// 	fmt.Println(a)
	// 	function1 = a.add
	// 	fmt.Println(a)
	// 	// 声明一个闭包并直接执行
	// 	// 此闭包返回值是另外一个闭包（带参闭包）
	// 	returnFunc := func() func(int, string) (int, string) {
	// 		fmt.Println("this is a anonymous function")
	// 		return func(i int, s string) (int, string) {
	// 			return i, s
	// 		}
	// 	}()

	// 	// 执行returnFunc闭包并传递参数
	// 	ret1, ret2 := returnFunc(1, "test")
	// 	fmt.Println("call closure function, return1 = ", ret1, "; return2 = ", ret2)

	//		fmt.Println("a.i = ", a.i)
	//		fmt.Println("after call function1, a.i = ", function1(1))
	//		fmt.Println("a.i = ", a.i)
	//	}
	// type A struct {
	// 	i int
	// }

	// // 定义方法
	// func (a *A) add(v int) int {
	// 	a.i += v
	// 	return a.i
	// }

	// // 声明函数变量
	// var function func(int) int

	// var sumAll, sumOdd, sumMultiple3 int
	// a, b := 0, 1
	// for n := 1; n <= 30; n++ {
	// 	current := a
	// 	sumAll += current
	// 	if current%2 != 0 {
	// 		sumOdd += current
	// 	}
	// 	if current%3 == 0 {
	// 		sumMultiple3 += current
	// 	}
	// 	a, b = b, a+b
	// }
	// fmt.Printf("前三十项的和: %d\n", sumAll)
	// fmt.Printf("前三十项中所有奇数项的和: %d\n", sumOdd)
	// fmt.Printf("前三十项中所有3的倍数项的和: %d\n", sumMultiple3)

	// const answer int = 32
	// var guess int
	// fmt.Println("您有100次机会，请猜测一个100以内的数字")
	// for attemts := 99; attemts >= 0; attemts-- {
	// 	fmt.Scanln(&guess)
	// 	if guess == answer {
	// 		fmt.Println("恭喜您，猜中了！")
	// 		return
	// 	} else if guess < answer {
	// 		fmt.Print("数字太小，您的剩余机会还剩：", attemts)
	// 	} else if guess > answer {
	// 		fmt.Println("数字太大，您的剩余机会还剩：", attemts)
	// 	}
	// 	if attemts == 0 {
	// 		fmt.Print("很遗憾，您没有猜中，正确答案是：", answer)
	// 	}

	// }
	//fmt.Println(fobona(3))
	//var guess string
	// fmt.Println("请依序输入年、月、日：")
	// year := "2008"
	// m := "5"
	// day := "21"
	// fmt.Println(stuer(year, m, day))
	// testCases := []int{1, 35, -9998877, 0, -100, 999}

	// for _, num := range testCases {
	// 	original := num
	// 	negate(&num) // 传递变量的地址给negate函数

	// 	fmt.Printf("原始值: %d, 取反后: %d\n", original, num)
	// }

	// nus := [6]int{3, 2, 4, 9, 7, 5}
	// sortArray(&nus)
	// fmt.Print(nus)
	// Student 结构体表示学生信息
	// type Student struct {
	// 	Name  string
	// 	Age   int
	// 	Score float64
	// }

	// var students []Student

	// // 添加一些示例学生数据
	// students = append(students, Student{"张三", 18, 85.5})
	// students = append(students, Student{"李四", 19, 92.0})
	// students = append(students, Student{"王五", 17, 78.5})

	// // 获取用户输入
	// var name string
	// var age int
	// var score float64

	// fmt.Println("请输入学生信息:")
	// fmt.Print("姓名: ")
	// fmt.Scanln(&name)
	// fmt.Print("年龄: ")
	// fmt.Scanln(&age)
	// fmt.Print("考试成绩: ")
	// fmt.Scanln(&score)

	// // 将新学生添加到列表中
	// students = append(students, Student{name, age, score})

	// // 输出学生成绩表
	// fmt.Println("\n============ 学生成绩表 ============")
	// fmt.Printf("%-10s %-6s %-8s\n", "姓名", "年龄", "成绩")
	// fmt.Println("----------------------------------")

	// for _, student := range students {
	// 	fmt.Printf("%-10s %-6d %-8.1f\n", student.Name, student.Age, student.Score)
	// }

	// fmt.Println("==================================")
	// var b int = 4
	// fmt.Println("local variable, b = ", b)
	// if b := 3; b == 3 {
	// 	fmt.Println("if statement, b = ", b)
	// 	b--
	// 	fmt.Println("if statement, b = ", b)
	// }
	// fmt.Println("local variable, b = ", b)

	// var a [5]int
	// fmt.Println("a = ", a)

	// var marr [2]map[string]string
	// fmt.Println("marr = ", marr)
	// // map的零值是nil，虽然打印出来是非空值，但真实的值是nil
	// //marr[0]["test"] = "1"

	// // 声明以及初始化
	// var b [5]int = [5]int{1, 2, 3, 4, 5}
	// fmt.Println("b = ", b)

	// // 类型推导声明方式
	// var c = [5]string{"c1", "c2", "c3", "c4", "c5"}
	// fmt.Println("c = ", c)

	// d := [3]int{3, 2, 1}
	// fmt.Println("d = ", d)

	// // 使用 ... 代替数组长度
	// autoLen := [...]string{"auto1", "auto2", "auto3"}
	// fmt.Println("autoLen = ", autoLen)

	// // 声明时初始化指定下标的元素值
	// positionInit := [5]string{1: "position1", 3: "position3", 2: "aa"}
	// fmt.Println("positionInit = ", positionInit)

	// // 初始化时，元素个数不能超过数组声明的长度
	// //overLen := [3]int{1, 2, 3}

	//a := [5]int{5, 4, 3, 2, 1}

	// // 方式1，使用下标读取数据
	// element := a[2]
	// fmt.Println("element = ", element)

	// // 方式2，使用range遍历
	// for i, v := range a {
	// 	fmt.Println("index = ", i, "value = ", v)
	// }

	// for i := range a {
	// 	fmt.Println("only index, index = ", i)
	// }

	// // 读取数组长度
	// fmt.Println("len(a) = ", len(a))
	// // 使用下标，for循环遍历数组
	// for i := 0; i < len(a); i++ {
	// 	fmt.Println("use len(), index = ", i, "value = ", a[i])
	// }

	// a := [3][2]int{
	// 	{0, 1},
	// 	{2, 3},
	// 	{4, 5},
	// }
	// fmt.Println("a = ", a)

	// // 三维数组
	// b := [3][2][2]int{
	// 	{{0, 1}, {2, 3}},
	// 	{{4, 5}, {6, 7}},
	// 	{{8, 9}, {10, 11}},
	// }
	// fmt.Println("b = ", b)

	// // 也可以省略各个位置的初始化,在后续代码中赋值
	// c := [3][3][3]int{}
	// c[2][2][1] = 5
	// c[1][2][1] = 4
	// fmt.Println("c = ", c)

	// a := [3][2][2]int{
	// 	{{0, 1}, {2, 3}},
	// 	{{4, 5}, {6, 7}},
	// 	{{8, 9}, {10, 11}},
	// }

	// layer1 := a[0]
	// layer2 := a[0][1]
	// element := a[0][1][1]
	// fmt.Println(layer1)
	// fmt.Println(layer2)
	// fmt.Println(element)

	// // 多维数组遍历时，需要使用嵌套for循环遍历
	// for i, v := range a {
	// 	fmt.Println("index = ", i, "value = ", v)
	// 	for j, inner := range v {
	// 		fmt.Println("inner, index = ", j, "value = ", inner)
	// 		for k, s := range inner {
	// 			fmt.Printf("inner, index = ", k, "value = ", s)
	// 		}
	// 	}
	// }
	// // 方式1，声明并初始化一个空的切片
	// var s1 []int = []int{}

	// // 方式2，类型推导，并初始化一个空的切片
	// var s2 = []int{}

	// // 方式3，与方式2等价
	// s3 := []int{}

	// // 方式4，与方式1、2、3 等价，可以在大括号中定义切片初始元素
	// s4 := []int{1, 2, 3, 4}

	// // 方式5，用make()函数创建切片，创建[]int类型的切片，指定切片初始长度为0
	// s5 := make([]int, 0)

	// // 方式6，用make()函数创建切片，创建[]int类型的切片，指定切片初始长度为2，指定容量参数4
	// s6 := make([]int, 2, 4)

	// // 方式7，引用一个数组，初始化切片
	// a := [5]int{6,5,4,3,2}
	// // 从数组下标2开始，直到数组的最后一个元素
	// s7 := arr[2:]
	// // 从数组下标1开始，直到数组下标3的元素，创建一个新的切片
	// s8 := arr[1:3]
	// // 从0到下标2的元素，创建一个新的切片
	// s9 := arr[:2]

	// a := [5]int{6, 5, 4, 3, 2}
	// // 从数组下标2开始，直到数组的最后一个元素
	// s7 := a[2:]
	// // 从数组下标1开始，直到数组下标3的元素，创建一个新的切片
	// s8 := a[1:3]
	// // 从0到下标2的元素，创建一个新的切片
	// s9 := a[:2]
	// fmt.Println(a)
	// fmt.Println(s7)
	// fmt.Println(s8)
	// fmt.Println(s9)
	// a[0] = 9
	// a[1] = 8
	// a[2] = 7
	// fmt.Println(a)
	// fmt.Println(s7)
	// fmt.Println(s8)
	// fmt.Println(s9)
	// s1 := []int{5, 4, 3, 2, 1}
	// // 下标访问切片
	// e1 := s1[0]
	// e2 := s1[1]
	// e3 := s1[2]
	// fmt.Println(s1)
	// fmt.Println(e1)
	// fmt.Println(e2)
	// fmt.Println(e3)

	// // 向指定位置赋值
	// s1[0] = 10
	// s1[1] = 9
	// s1[2] = 8
	// fmt.Println(s1)

	// // range迭代访问切片
	// for i, v := range s1 {
	// 	fmt.Println("before modify, s1[%d] = %d", i, v)
	// }

	// var nilSlice []int
	// fmt.Println("nilSlice length:", len(nilSlice))
	// fmt.Println("nilSlice capacity:", len(nilSlice))

	// s2 := []int{9, 8, 7, 6, 5}
	// fmt.Println("s2 length: ", len(s2))
	// fmt.Println("s2 capacity: ", cap(s2))

	// s3 := []int{}
	// fmt.Println("s3 = ", s3)

	// // append函数追加元素
	// s3 = append(s3)
	// s3 = append(s3, 1)
	// s3 = append(s3, 2, 3)
	// fmt.Println("s3 = ", s3)

	// s4 := []int{1, 2, 4, 5}
	// s4 = append(s4[:2], append([]int{3}, s4[2:]...)...)
	// fmt.Println("s4 = ", s4)

	// s5 := []int{1, 2, 3, 5, 4}
	// //s5 = append(s5[:3], s5[4:]...)
	// s5 = append(s5[:1], s5[2:]...)
	// fmt.Println("s5 = ", s5)

	// src1 := []int{1, 2, 3}
	// dst1 := make([]int, 4, 5)

	// src2 := []int{1, 2, 3, 4, 5}
	// dst2 := make([]int, 3, 3)

	// fmt.Println("before copy, src1 = ", src1)
	// fmt.Println("before copy, dst1 = ", dst1)

	// fmt.Println("before copy, src2 = ", src2)
	// fmt.Println("before copy, dst2 = ", dst2)

	// copy(dst1, src1)
	// copy(dst2, src2)

	// fmt.Println("before copy, src1 = ", src1)
	// fmt.Println("before copy, dst1 = ", dst1)

	// fmt.Println("before copy, src2 = ", src2)
	// fmt.Println("before copy, dst2 = ", dst2)
	// var m1 map[string]string
	// fmt.Println("m1 length:", len(m1))

	// m2 := make(map[string]string)
	// fmt.Println("m2 length:", len(m2))
	// fmt.Println("m2 =", m2)

	// m3 := make(map[string]string, 10)
	// fmt.Println("m3 length:", len(m3))
	// fmt.Println("m3 =", m3)

	// m4 := map[string]string{}
	// fmt.Println("m4 length:", len(m4))
	// fmt.Println("m4 =", m4)

	// m5 := map[string]string{
	// 	"key1": "value1",
	// 	"key2": "value2",
	// }
	// fmt.Println("m5 length:", len(m5))
	// fmt.Println("m5 =", m5)

	// m := make(map[string]int, 10)

	// m["1"] = int(1)
	// m["2"] = int(2)
	// m["3"] = int(3)
	// m["4"] = int(4)
	// m["5"] = int(5)
	// m["6"] = int(6)

	// // 获取元素
	// //fmt.Println(m)
	// // value1 := m["1"]
	// // fmt.Println("m[\"1\"] =", value1)

	// // value1, exist := m["5"]
	// // fmt.Println("m[\"5\"] =", value1, ", exist =", exist)

	// // valueUnexist, exist := m["10"]
	// // fmt.Println("m[\"10\"] =", valueUnexist, ", exist =", exist)

	// // 修改值
	// // fmt.Println("before modify, m[\"2\"] =", m["2"])
	// m["2"] = 20
	// // fmt.Println(m)
	// // fmt.Println("after modify, m[\"2\"] =", m["2"])

	// // // 获取map的长度
	// // fmt.Println("before add, len(m) =", len(m))
	// m["10"] = 10
	// m["11"] = 12
	// // fmt.Println(m)
	// // fmt.Println("after add, len(m) =", len(m))

	// // 遍历map集合main
	// // for key, value := range m {
	// // 	fmt.Println("iterate map, m[", key, "] =", value)
	// // }

	// // 使用内置函数删除指定的key
	// _, exist_10 := m["10"]
	// fmt.Println(m)
	// fmt.Println(exist_10)
	// fmt.Println("before delete, exist 10: ", exist_10)
	// delete(m, "10")
	// _, exist_10 = m["7"]
	// fmt.Println(m)
	// fmt.Println("after delete, exist 10: ", exist_10)

	// 在遍历时，删除map中的key
	// for key := range m {
	// 	fmt.Println("iterate map, will delete key:", key)
	// 	delete(m, key)
	// }
	// fmt.Println("m = ", m)

	// 	m := make(map[string]int)
	// 	fmt.Println(m)
	// 	m["a"] = 1
	// 	m["c"] = 5
	// 	fmt.Println(m)
	// 	receiveMap(m)
	// 	fmt.Println("m =", m)
	// }

	// func receiveMap(param map[string]int) {
	// 	fmt.Println("before modify, in receiveMap func: param[\"a\"] = ", param["a"])
	// 	param["a"] = 2
	// 	param["b"] = 3
	// m := make(map[string]int)

	// go func() {
	// 	for {
	// 		m["a"]++
	// 	}
	// }()

	// go func() {
	// 	for {
	// 		m["a"]++
	// 		fmt.Println(m["a"])
	// 	}
	// }()

	// select {
	// case <-time.After(time.Second * 5):
	// 	fmt.Println("timeout, stopping")
	// }

	// m := make(map[string]int)
	// var wg sync.WaitGroup
	// var lock sync.Mutex
	// wg.Add(2)

	// go func() {
	// 	for {
	// 		lock.Lock()
	// 		m["a"]++
	// 		lock.Unlock()
	// 	}
	// }()

	// go func() {
	// 	for {
	// 		lock.Lock()
	// 		m["a"]++
	// 		fmt.Println(m["a"])
	// 		lock.Unlock()
	// 	}
	// }()

	// select {
	// case <-time.After(time.Second * 5):
	// 	fmt.Println("timeout, stopping")
	// }
	// str1 := "abc123"
	// for index := range str1 {
	// 	fmt.Printf("str1 -- index:%d, value:%d\n", index, str1[index])
	// }

	// str2 := "测试中文"
	// for index := range str2 {
	// 	fmt.Printf("str2 -- index:%d, value:%d\n", index, str2[index])
	// }
	// fmt.Printf("len(str2) = %d\n", len(str2))

	// runesFromStr2 := []rune(str2)
	// bytesFromStr2 := []byte(str2)
	// fmt.Printf("len(runesFromStr2) = %d\n", len(runesFromStr2))
	// fmt.Printf("len(bytesFromStr2) = %d\n", len(bytesFromStr2))

	// array := [...]int{1, 2, 3}
	// slice := []int{4, 5, 6}

	// // 方法1：只拿到数组的下标索引
	// for index := range array {
	// 	fmt.Printf("array -- index=%d value=%d \n", index, array[index])
	// }
	// for index := range slice {
	// 	fmt.Printf("slice -- index=%d value=%d \n", index, slice[index])
	// }
	// fmt.Println()

	// // 方法2：同时拿到数组的下标索引和对应的值
	// for index, value := range array {
	// 	fmt.Printf("array -- index=%d index value=%d \n", index, array[index])
	// 	fmt.Printf("array -- index=%d range value=%d \n", index, value)
	// }
	// for index, value := range slice {
	// 	fmt.Printf("slice -- index=%d index value=%d \n", index, slice[index])
	// 	fmt.Printf("slice -- index=%d range value=%d \n", index, value)
	// }
	fmt.Println()

}

// func fobona(n int) int {
// 	if n == 1 {
// 		return 0
// 	} else if n == 2 {
// 		return 1
// 	}
// 	return fobona(n-1) + fobona(n-2)
// }

// func stuer(year, moth, day string) string {
// 	str := "[" + year + "]年[" + moth + "]" + "月[" + day + "]日"
// 	return str

// }
// func negate(num *int) {
// 	*num = -*num
// }
// func paixunum(nums []int) []int {
// 	seen := make(map[int]bool)
// 	rest := []int{}
// 	for _, num := range nums {
// 		seen[num] = true
// 		rest = append(rest, num)

// 	}
// 	return rest
// }
// func sortArray(arr *[6]int) {
// 	// 使用冒泡排序算法
// 	n := len(arr)
// 	for i := 0; i < n-1; i++ {
// 		for j := 0; j < n-i-1; j++ {
// 			if arr[j] > arr[j+1] {
// 				// 交换元素
// 				arr[j], arr[j+1] = arr[j+1], arr[j]
// 			}
// 		}
// 	}

// }

// func sortArray1(arr *[6]int) {
// 	// 将数组转换为切片，使用标准库排序，然后复制回数组
// 	slice := arr[:]
// 	sort.Ints(slice)
// }

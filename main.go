package main

import "fmt"

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

	var sumAll, sumOdd, sumMultiple3 int
	a, b := 0, 1
	for n := 1; n <= 30; n++ {
		current := a
		sumAll += current
		if current%2 != 0 {
			sumOdd += current
		}
		if current%3 == 0 {
			sumMultiple3 += current
		}
		a, b = b, a+b
	}
	fmt.Printf("前三十项的和: %d\n", sumAll)
	fmt.Printf("前三十项中所有奇数项的和: %d\n", sumOdd)
	fmt.Printf("前三十项中所有3的倍数项的和: %d\n", sumMultiple3)

}

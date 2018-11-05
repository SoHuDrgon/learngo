package main

import "fmt"

//函数function
//go函数不支持嵌套，重载和默认参数
//支持以下特性
//无需声明原型，不定长度变参，多返回值，命名返回值参数，匿名参数，闭包
//定义函数使用关键字func，但左大括号不能另起一行
//函数可以作为一种类型使用
//func A(a int, b string) {int, string} {
//func A(a, b, c int) {
//	a, b, c = 1, 2, 3
//return a, b ,c
//}
func main() {
	a, b := 1, 2
	A(a, b)
	fmt.Println(a, b)
	s2 := []int{1, 2, 3, 4}
	B(s2)
	fmt.Println(s2)
	a1 := 1
	C(&a1)
	fmt.Println(a1)
	a2 := D
	a2()
	//匿名函数
	a3 := func() {
		fmt.Println("Func E")
	}
	a3()

	f := closure(10)
	fmt.Println(f(1))
	fmt.Println(f(2))

	//fmt.Println("a")
	//defer fmt.Println("b")
	//defer fmt.Println("c")
	//a,c,b
	//作为参数传递进去的
	//for i := 0; i < 3; i++ {
	//	defer fmt.Println(i)
	//}
	//2,1,0
	//下面是作为闭包，i作为地址引用，引用的局部变量，退出循环体的时候i变成了3
	//for i := 0; i < 3; i++ {
	//	defer func() {
	//		fmt.Println(i)
	//	}()
	//}
	E()
	F()
	G()

	var fs = [4]func(){}
	for i := 0; i < 4; i++ {
		//这里是将i作为了一个参数，遵循规则是值拷贝和传递，defer定义逆序输出
		defer fmt.Println("defer i = ", i)
		//通过defer定义的，地址是先进后出，后进先出，也是4
		defer func() { fmt.Println("defer_closure i = ", i) }()
		//匿名函数的i是从外面夺过来的，是引用地址，输出的地址是4
		fs[i] = func() { fmt.Println("closuer i = ", i) }
	}

	for _, f := range fs {
		f()
	}
}

func A(s ...int) {
	s[0] = 3
	s[1] = 4
	fmt.Println(s)
}

//拷贝值
func B(s1 []int) {
	s1[0] = 5
	s1[1] = 6
	s1[2] = 7
	s1[3] = 8
	fmt.Println(s1)
}

//拷贝地址
func C(a1 *int) {
	*a1 = 2
	fmt.Println(*a1)
}

//类型
func D() {
	fmt.Println("Func D")
}

//闭包
func closure(x int) func(int) int {
	fmt.Printf("%p\n", &x)
	return func(y int) int {
		fmt.Printf("%p\n", &x)
		return x + y
	}
}

//defer
//执行方式类似其他语言中的析构函数，在函数执行结束后安装调用顺序的相反顺序逐个执行
//即使函数发生严重错误也会执行
//支持匿名函数调用
//常用于资源清理，文件关闭，解锁及记录时间等操作
//通过与匿名函数配合可以在return后修改函数计算结果
//如果函数体内某个变量作为defer时匿名函数的参数，则在定义defer时已经获得了拷贝，否则则是引用某个变量的地址
//Go没有异常机制，但是有panic/recover模式来处理错误
//Panic可以在任何地方引发，但recover只有在defer调用的函数中有效

func E() {
	fmt.Println("Func E")
}
func F() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recover in F")
		}
	}()
	panic("Panic in F")
}
func G() {
	fmt.Println("Func G")
}

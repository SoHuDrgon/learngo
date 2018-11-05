package main

import (
	"fmt"
	"strconv"
)

type (
//byte int8
//rune int32
//ByteSize int64
)

func main() {
	//var b int32 = 321
	//fmt.Println(b)
	c, d, e, f := 1, 2, 3, 4
	fmt.Println(c, e, f) //应用在函数返回值时
	fmt.Println(c, d, e, f)
	//变量的类型转换
	var a float32 = 100.1
	fmt.Println(a)
	b := int(a)
	fmt.Println(b)
	//var g bool = true
	//fmt.Println(g)
	//h := int(g) //bool和int不存在兼容
	//fmt.Println(h)
	var g int = 64
	//h := string(g)
	h := strconv.Itoa(g)
	fmt.Println(h)
	g, _ = strconv.Atoi(h)
	fmt.Println(g)
	//常量的定义
	//常量的值在编译的时候就已经确定
	//常量的定义格式与变量基本相同
	//等号右侧必须是常量或者常量表达式
	//常量表达式中的函数必须是内置函数
	const cla int = 1
	const clb = 10
	const (
		text   = "123"
		length = len(text)
		num    = clb * 20
	)
	fmt.Println(cla)
	fmt.Println(clb)
	fmt.Println(text, length, num)
	//同时定义多个常量
	const i, j, k = 1, "2", 3
	const (
		text2, length2, num2 = "456", len(text), k * 10
	)
	fmt.Println(i, j, k)
	fmt.Println(text2, length2, num2)
	//常量的初始化规律与枚举
	//在定义常量组时，如果不提供初始值，则表示将使用上行的表达式
	//使用相同的表达式不代表具有相同的值
	//iota是常量的计数器，从0开始，组中每定义一个常量自动递增1
	//通过初始化规则与iota可以达到枚举的效果
	//每遇到一个const关键字，iota就会重置为0
	const (
		a1 = 1
		b1
		c1
	)
	fmt.Println(a1, b1, c1)

	const (
		a2 = 'A'
		b2
		c2 = iota
		d2
	)
	fmt.Println(a2, b2, c2, d2)
	const (
		e2 = iota
	)
	fmt.Println(e2)
	const (
		MAX_COUNT  = 'A'
		_MAX_COUNT = 23
		cMAX_COUNT = 32
	)
	fmt.Println(MAX_COUNT, _MAX_COUNT, cMAX_COUNT)
	fmt.Println(^2)
	fmt.Println(1 ^ 2)
	fmt.Println(!true)
	fmt.Println(5 * 2)
	fmt.Println(5 / 2)
	fmt.Println(5 % 2)
	fmt.Println(1 << 10 << 10 >> 10)
	fmt.Println('a' << 10)
	/*
		6:	0110
		11:	1011
		----------
		& 	0010
		|	1111
		^	1101
		&^	0100
	*/
	fmt.Println(6 & 11)
	fmt.Println(6 | 11)
	fmt.Println(6 ^ 11)
	fmt.Println(6 &^ 11)

	a3 := 1
	if a3 > 0 && (10/a3) > 1 {
		fmt.Println("Ok")
	}
	const (
		_          = iota
		KB float64 = 1 << (iota * 10)
		MB
		GB
		TB
		PB
		EB
		ZB
		YB
	)
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(TB)
	fmt.Println(PB)
	fmt.Println(EB)
	fmt.Println(ZB)
	fmt.Println(YB)
	//指针
	//Go虽然保留了指针，但与其它编程语言不同的是，在Go当中不支持指针运算以及" -> " 运算符，而直接采用
	//"."选择符来操作指针对象的成员
	//-- 操作符"&" 取变量地址，使用"*" 通过指针间接访问目标对象
	//-- 默认值为nil而非NULL
	//递增递减语句
	//在Go当中， ++ 与 -- 是作为语句而并不是作为表达式
	a4 := 1024
	a4++
	var p *int = &a4
	fmt.Println(*p)
	a4--
	fmt.Println(a4)

	//判断语句if
	//条件表达式没有括号
	//支持一个初始化表达式（可以是并行方式）
	//左大括号必须和条件语句或else在同一行
	//支持单行模式
	//初始化语句中的变量为block级别，同事隐藏外部同名变量
	a5 := true
	if a5, b5, c5 := 3, 2, 3; a5+b5+c5 > 6 {
		fmt.Println("大于6")
	} else {
		fmt.Println("小于等于6")
		fmt.Println(a5)
	}
	fmt.Println(a5)

	//循环语句for
	//Go只有for一个循环语句关键字，但支持3种形式
	//初始化和步进表达式可以是多个值
	//条件语句每次循环都会被重新检查，因此不建议在条件语句中使用函数，尽量提前计算好条件并以变量或常量代替
	//左大括号必须和条件语句在同一行

	a6 := 1
	for {
		a6++
		if a6 > 3 {
			break
		}
		fmt.Println(a6)
	}

	a7 := 1
	for a7 <= 3 {
		a7++
	}
	fmt.Println(a7)

	a8 := 1
	for i := 0; i < 3; i++ {
		a8++
	}
	fmt.Println(a8)

	//a9 := "string"
	//l := len(a9)
	//for i := 0; i < l; i++ {
	//	a9++
	//	fmt.Println(a9)
	//}
	//fmt.Println("Over")

	//选择语句switch
	//可以使用任何类型或表达式作为条件语句
	//不需要写break，一旦条件符合自动终止
	//如果希望继续执行下一个case，需要使用fallthrough语句
	//支持一个初始化表达式(可以是并行方式),右侧需跟分号
	//左大括号必须和条件语句在同一行

	b3 := 0
	switch b3 {
	case 0:
		fmt.Println("b3=0")
	case 1:
		fmt.Println("b3=1")
	}
	fmt.Println(b3)

	b4 := 1
	switch {
	case b4 >= 0:
		fmt.Println("b4=0")
		fallthrough
	case b4 >= 1:
		fmt.Println("b4=1")
	}
	fmt.Println(b4)

	switch b5 := 1; {
	case b5 >= 0:
		fmt.Println("b5=0")
		fallthrough
	case b5 >= 1:
		fmt.Println("b5=1")
	}

	//跳转语句goto， break， continue
	//三个语法都可以配合标签使用
	//标签区分大小写，若不使用会造成编译错误
	//Break与continue配合标签可以用于多层循环的跳出
	//got是调整执行位置，与其他两个语句配合标签的结果并不相同
	fmt.Println("break example")
LABEL:
	for {
		for i := 0; i < 10; i++ {
			if i > 2 {
				break LABEL
			} else {
				fmt.Println(i)
			}
		}
	}
	fmt.Println("continue example")
LABEL1:
	for i1 := 0; i1 < 10; i1++ {
		for {
			fmt.Println(i1)
			continue LABEL1
			goto LABEL1 //死循环,调整程序执行位置到标签那里
		}
	}
}

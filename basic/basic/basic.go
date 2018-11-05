package basic

import (
	"fmt"
	"math"
	"math/cmplx"
)

//作用域是包内部的变量。
var aa = 3
var ss = "kkk"

// bb := true
var bb = true

//简写
var (
	a1 = 1
	b1 = false
	c1 = true
)

//内建变量类型
//bool,string
//(u)int,(u)int8,(u)int16,(u)int32,(u)int64,uintptr
//byte,rune  char 1字节类型坑多，所以用rune替代
//float32,float64,complex64,complex128	浮点数 复数
func variableZeroValue() {
	var a int
	var s string
	fmt.Println(a, s)
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

//省略int，string
func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

//使用:= 定义变量
func variableShorter() {
	a, b, c, s := 4, 6, true, "def"
	b = 5
	fmt.Println(a, b, c, s)
}

//欧拉公式实现
func euler() {
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
	fmt.Println(cmplx.Exp(1i*math.Pi) + 1)
	//显示小数点后3位
	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)
}

//强制类型转换
//类型转换时强制的
//var a,b int = 3，4
//c = math.Sqrt(float64(a * a + b * b)) 不能隐式转换
//必须强制转换
//c = int(math.Sqrt(float64(a * a + b * b)))
func triangle() {
	var a, b int = 3, 4
	//var c int
	////c = math.Sqrt(float64(a * a + b * b)) 不能隐式转换
	////必须强制转换
	//c = int(math.Sqrt(float64(a * a + b * b)))
	//fmt.Println(c)
	fmt.Println(calcTriangle(a, b))
}

func calcTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}

//常量的定义
//常量的数值可以作为各种类型使用
func consts() {
	//const filename string = "abc.txt"
	//const a, b = 3, 4
	const (
		filename = "abc.txt"
		a, b     = 3, 4
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

//常量定义枚举类型
//变量类型卸载变量名之后
//编译器可推测变量类型
//没有char， 没有rune
//原生支持复数类型
func enums() {
	const (
		cpp    = 0
		java   = 1
		python = 2
		golang = 3
	)
	fmt.Println(cpp, java, python, golang)
}

//iota实现自增值枚举类型
func enums1() {
	const (
		cpp = iota
		_
		python
		golang
		javascript
	)
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp, javascript, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	fmt.Println("Hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	euler()
	triangle()
	consts()
	enums()
	enums1()
}

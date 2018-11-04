package main

import "fmt"

func adder() func(int) int {
	sum := 0 //需要一个自由变量在里面，会被改变
	return func(v int) int {
		sum += v
		return sum
	}
}

//正统函数式编程写法
type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder2(base + v)
	}
}
func main() {
	a := adder()
	for i := 0; i < 10; i++ {
		fmt.Println(a(i))
		fmt.Printf("0 + 1 + ... + %d= %d\n",
			i, a(i))
	}
	b := adder2(8)
	for i := 0; i < 10; i++ {
		var s int
		s, b = b(i)
		fmt.Printf("0 + 1 + ... + %d= %d\n",
			i, s)
	}
}

package main

import "fmt"

func main() {
	//数组array
	//定义数组的格式：var <varName> [n]<type>, n >= 0
	//数组长度也是类型的一部分，因此具有不同长度的数组为不同类型
	//注意区分指向数组的指针和指针数组
	//数组在Go中为值类型
	//数组之间可以使用==或!=进行比较，但不可以使用< or >
	//可以使用new来创建数组，此方法返回一个指向数组的指针
	//Go支持多维数组
	var a [2]int //如果是string，输出是空字符串
	var b [2]int
	b = a
	fmt.Println(b)
	a1 := [2]int{1, 1}
	fmt.Println(a1)
	a2 := [20]int{19: 1} //赋值给第二十个元素为1
	fmt.Println(a2)
	a3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(a3)
	a4 := [...]int{0: 1, 1: 2, 2: 3}
	fmt.Println(a4)
	a5 := [...]int{19: 1}
	var p *[20]int = &a5
	fmt.Println(p)
	x, y := 1, 2
	a6 := [...]*int{&x, &y}
	fmt.Println(a6)
	a7 := [2]int{1, 2}
	b2 := [2]int{2, 3}
	fmt.Println(a7 == b2)
	a8 := [10]int{}
	a8[1] = 2
	fmt.Println(a8)
	p1 := new([10]int)
	p1[1] = 2
	fmt.Println(p1)
	a9 := [2][3]int{
		{1, 1, 1},
		{2, 2, 2}}
	fmt.Println(a9)
	b3 := [...][3]int{
		{3, 4},
		{5, 6, 7}}
	fmt.Println(b3)
	//冒泡排序
	b4 := [...]int{5, 2, 6, 3, 9}
	fmt.Println(b4)
	num := len(b4)
	for i := 0; i < num; i++ {
		for j := i + 1; j < num; j++ {
			if b4[i] < b4[j] {
				temp := b4[i]
				b4[i] = b4[j]
				b4[j] = temp
			}
		}
	}
	fmt.Println(b4)
}

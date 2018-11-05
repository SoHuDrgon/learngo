package main

import "fmt"

//切片Slice
//其本身不是数组，它指向底层的数组
//作为变长数组的替代方案，可以关联底层数组的局部或全部
//为引用类型
//可以直接创建或从底层数组获取生成
//使用len()获取元素个数，cap()获取容量
//一般使用make()创建
//如果多个slice指向相同底层数组，其中一个的值改变会影响全部
//make([], len, cap)
//其中cap可以省略，则和len的值相同
//len表示存数的元素个数，cap表示容量

func main() {
	var s1 []int
	fmt.Println(s1)
	//a := [10]int{}
	//fmt.Println(a)
	//s2 := a[9]
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//s2 := a[5:]
	s2 := a[5:10]
	fmt.Println(s2)
	s3 := make([]int, 3, 10) //10表示分配的内存地址，连续的，不超过10个就固定分配10个地址使用
	//超过10个的话，go会分配20个内存地址，翻倍分配，预先分配好效率高
	fmt.Println(s3, len(s3), cap(s3))
	a1 := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k'}
	s4 := a1[2:5]
	fmt.Println(string(s4))
	s5 := a1[3:5]
	fmt.Println(string(s5))
	//Reslice
	//reslice是索引以被slice的切片为准
	//索引不可以超过被slice的切片的容量cap()值
	//索引越界不会导致底层数组的重新分配而是引发错误
	sa := a1[2:5]
	sb := sa[1:3]
	fmt.Println(string(sb))
	fmt.Println(len(sa), cap(sa))
	sc := sa[3:5]
	fmt.Println(string(sc))
	//Append
	//可以在slice尾部追加元素
	//可以将一根slice追加在另一个slice尾部
	//如果最终长度未超过追加到slice的容量则返回原始slice
	//如果超过追加到的slice的容量则将重新分配数组并拷贝原始数据
	s6 := make([]int, 3, 6)
	fmt.Printf("%p\n", s6)
	s6 = append(s6, 1, 2, 3)
	fmt.Printf("%v %p\n", s6, s6)
	s6 = append(s6, 1, 2, 3) //重新分配内存地址
	fmt.Printf("%v %p\n", s6, s6)

	a2 := []int{1, 2, 3, 4, 5}
	s7 := a2[2:5]
	s8 := a2[1:3]
	fmt.Println(s7, s8)
	//s7[0] = 9 //都会更改
	//fmt.Println(s7, s8)
	s9 := a2[1:3]
	s9 = append(s9, 1, 2, 3, 1, 10, 11, 12)
	s7[0] = 9
	fmt.Println(s7, s9)

	//Copy
	a3 := []int{1, 2, 3, 4, 5, 6}
	//a4 := []int{7, 8, 9}
	//copy(a3, a4)
	//fmt.Println(a3) //[7 8 9 4 5 6]
	//copy(a4, a3)
	//fmt.Println(a4) //[1 2 3]
	//copy(a3[2:4], a4[1:3])
	//fmt.Println(a4)
	a5 := s3[:]
	fmt.Println(a3)
	fmt.Println(a5)
}

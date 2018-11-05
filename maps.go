package main

import (
	"fmt"
	"sort"
)

//map
//类似其他语言中的哈希表或者字典，以key-value形式存储数据
//key必须支持==或!=比较运算的类型，不可以是函数，map或slice
//map查找比线性搜索快很多，但比使用索引访问数据的类型慢100倍
//map使用make()创建，支持:=这种简写方式
//make([keyType]valueType, cap), cap表示容量，可省略
//超出容量时会自动扩容，但尽量提供一个合理的初始值
//使用len()获取元素个数
//键值对不存在时自动添加，使用delete()删除某键值对
//使用for range对map和slice进行迭代操作

func main() {
	//map的创建
	var m map[int]string
	m = map[int]string{}
	fmt.Println(m)

	var m1 map[int]string
	m1 = make(map[int]string)
	fmt.Println(m1)

	m2 := make(map[int]string)
	m2[1] = "Ok"
	fmt.Println(m2)
	a := m2[1]
	fmt.Println(a)

	var m3 map[int]map[int]string
	m3 = make(map[int]map[int]string)
	m3[1] = make(map[int]string)
	m3[1][1] = "Good"
	a1 := m3[1][1]
	fmt.Println(a1)

	var m4 map[int]map[int]string
	m4 = make(map[int]map[int]string)
	a2, ok := m4[2][1]
	if !ok {
		m4[2] = make(map[int]string)
	}
	m4[2][1] = "fuck you"
	a2, ok = m4[2][1]
	fmt.Println(a2, ok)

	sm := make([]map[int]string, 5)
	for _, v := range sm { //v是一个拷贝，对v的任何操作都不会影响sm本身
		v = make(map[int]string, 1)
		v[1] = "Maps"
		fmt.Println(v)
	}
	fmt.Println(sm) //sm的map为空
	//
	sm1 := make([]map[int]string, 5)
	for i := range sm {
		sm1[i] = make(map[int]string, 1)
		sm1[i][1] = "Maps"
		fmt.Println(sm1[i])
	}
	fmt.Println(sm1)

	m5 := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"}
	s := make([]int, len(m5))
	i := 0
	for k, _ := range m5 {
		s[i] = k
		i++
	}
	sort.Ints(s)
	fmt.Println(s)

	m6 := map[int]string{1: "a", 2: "b", 3: "c"}
	//m7 := map[string]int{"a":1, "b":2, "c":3}
	fmt.Println(m6)
	m7 := make(map[string]int)
	for k, v := range m6 {
		m7[v] = k
	}
	fmt.Println(m7)
}

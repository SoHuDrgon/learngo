package main

import (
	"fmt"
	"reflect"
)

//定义一个用户结构体
type User1 struct {
	Id   int
	Name string
	Age  int
}

type Manager struct {
	User1 //定义了一个匿名引用
	title string
}

func main() {
	m := Manager{User1: User1{1, "OK", 15}, title: "123"}
	t := reflect.TypeOf(m)

	//取得类型中的字段是否为匿名字段
	fmt.Printf("%6v\n", t.Field(0))
	/**
	打印内容：{User main.User  0 [ 0] true}，其中true表示是匿名类型
	那么想要取匿名类型中的字段又该怎么取呢？这里需要使用序号组，传入要取的切片即可
	*/
	fmt.Printf("%v\n", t.FieldByIndex([]int{0, 0}))
	/**
	其中上面切片传入的是{0, 0},
	1. 第一个0表示当前结构Manager取匿名User是第一个即为0
	2. 第二个0表示取得的结构User中要取第一个元素Id相对于User来说也是第一个即为0，如果要取Name则需传入[]int{0, 1}
	那么既然可以取出来内容，那么我们就可以尝试着进行修改，怎么做呢？
	*/
	tchage := reflect.ValueOf(&m)                       //想要修改和我们之前所说的传入值类型和指针类型是一致的，要想修改需要传入对应指针类型
	tchage.Elem().FieldByIndex([]int{0, 0}).SetInt(999) //传入指针需要通过 .Elem() 来取得对应的值内容，之后再想取哪个再继续使用序号组
	fmt.Println(tchage.Elem().FieldByName("title"))
	fmt.Println(tchage)
}

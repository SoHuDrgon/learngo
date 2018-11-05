package main

import (
	"fmt"
	"reflect"
)

//定义一个用户结构体
type User struct {
	Id   int
	Name string
	Age  int
}

//为接口绑定方法
func (u User) Hello() {
	fmt.Println("Hello World.")
}

//定义一个可接受任何类型的函数（空接口的使用规则）
func Info(o interface{}) {
	t := reflect.TypeOf(o)         //获取接受到到接口到类型
	fmt.Println("Type:", t.Name()) //打印对应类型到名称(这是reflect中自带到)

	//Kind()方法是得到传入类型到返回类型；下面执行判断传入类型是否为一个结构体
	if k := t.Kind(); k != reflect.Struct {
		fmt.Println("传入的类型有误，请检查!")
		return
	}

	v := reflect.ValueOf(o) //获取接受到到接口类型包含到内容(即其中到属性字段和方法)
	fmt.Println("Fields:")  //如何将其中到所有字段和内容打印出来呢？
	/**
	通过接口类型.NumField 获取当前类型所有字段个数
	*/
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)               //取得对应索引的字段
		val := v.Field(i).Interface() //取得当前字段对应的内容
		fmt.Printf("%6s: %v = %v\n", f.Name, f.Type, val)
	}
	/**
	通过接口类型.NumMethod 获取当前类型所有方法的个数
	*/
	fmt.Println("Method:")
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i) //取得对应索引的方法
		fmt.Printf("%6s: %v\n", m.Name, m.Type)
	}
}

func main() {
	u := User{1, "OK", 12}
	Info(u)
	//Info(&u) 如果传入的是结构体的地址或指针(pointer-interface)，那么在Info函数中的Kind方法进行判断时就会被拦截返回
}

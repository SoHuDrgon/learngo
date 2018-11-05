package main

import (
	"fmt"
	"reflect"
)

//定义一个用户结构体
type User2 struct {
	Id   int
	Name string
	Age  int
}

//为User绑定方法
func (u User2) HelloDisplay(name string) {
	fmt.Println("Hello", name, " my name is ", u.Name)
}

func main() {
	u := User2{1, "OK", 29}
	u.HelloDisplay("jack") //正常调用

	/**
	以下方式为反射调用，最优到代码写法就是新写一个方法且在开始是通过kind判断类型是否正确且需要判断有没有对应方法等
	*/
	v := reflect.ValueOf(u)                          //通过反射得到类型内容
	methodV := v.MethodByName("HelloDisplay")        //通过方法名称得道方法实体
	args := []reflect.Value{reflect.ValueOf("jack")} //设置反射传入的参数
	methodV.Call(args)
}

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

func main() {
	u := User{1, "OK", 13}
	fmt.Println(u)
	Set(&u)
	fmt.Println(u)
}

//定义一个可以接受任何类型的空接口
func Set(o interface{}) {
	v := reflect.ValueOf(o)
	//通过反射修改类型中的内容需要传入指针，为了防止传入有误故在这里进行相关过滤验证判断(这前这快是已经说过的)
	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		//reflect.Ptr对应为指针类型；v.Elem().CanSet()取得对应地址下的内容并查看其是否可以进行修改
		fmt.Println("传入的类型有误,请检查!")
		return
	} else {
		v = v.Elem() //将实际对象(包含详情内容)进行赋值
	}

	f := v.FieldByName("Name")
	f1 := v.FieldByName("Id1")
	if !f.IsValid() { //判断通过名称获取得到到内容是否为空值
		fmt.Println("没有Name对应属性字段")
		return
	}
	if !f1.IsValid() {
		fmt.Println("没有Id1对应属性字段")
	}
	if f.Kind() == reflect.String {
		f.SetString("HelloWorld")
	}
}

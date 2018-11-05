package main

import (
	"fmt"
)

/**
定义接口
*/
type USB interface {
	Name() string // 定义了名为Name的方法且返回类型为String
	Connecter
}

type Connecter interface {
	Connect()
}

/**
定义结构体
*/
type PhoneConnecter struct {
	// 定义内部变量
	name string
}

/**
真对结构体绑定的方法method
*/
func (phone PhoneConnecter) Name() string {
	return phone.name
}

func (phone PhoneConnecter) Connect() {
	fmt.Println("Connect:", phone.name)
}

func main() {
	a := PhoneConnecter{"PhoneConnecter"}
	a.Connect()
	Disconnect(a)
	DisconnectAll(a)
	DisconnectSwitch(a)

	// 演示接口类型转换
	phone := PhoneConnecter{"PhoneConnecter"} // 声明 PhoneConnecter，其实它是实现来 USB 接口的
	var b Connecter                           // 声明 Connecter 接口
	b = Connecter(phone)                      // 将 USB接口的类型强制转换为 Connecter 类型。即Go语言中接口转换只能从高——>低
	b.Connect()
}

func Disconnect(usb USB) {
	/**
	这里是通过类型断言来判定当前是哪个对象调用并打印
	类型断言格式：usb.(要判断的对象)，判断传入的是否为PhoneConnecter类型
	*/
	if phone, ok := usb.(PhoneConnecter); ok {
		fmt.Println("Disconnected：", phone.name)
		return
	}
	fmt.Println("Unknown decive.")
}

/**
Go语言中类型定义只要符合定义的接口那么它就实现来该接口，在java语言中都有一个顶级父类叫 Object。那么在Go语言中有吗？
当然有的，通过接口概念可以了解到当我们定义一个空接口的时候，任何类型都会继承它。那么针对Disconnect方法我们就可以定义一种更加广泛的应用方式
*/
func DisconnectAll(usb interface{}) {
	if phone, ok := usb.(PhoneConnecter); ok {
		fmt.Println("DisconnectAll：", phone.name)
		return
	}
	fmt.Println("Unknown decive.")
}

/**
在Go语言中可以用另外一种方式进行断言，叫：type switch
*/
func DisconnectSwitch(usb interface{}) {
	switch v := usb.(type) {
	case PhoneConnecter:
		fmt.Println("DisconnectSwitch：", v.name)
	default:
		fmt.Println("Unknown decive.")
	}
}

package main

import (
	"fmt"
)

func main() {
	//声明创建一个通道，存储类型为bool型
	c := make(chan bool)
	//启用一个goroutine,使用的是匿名方法方式
	go func() {
		fmt.Println("Go Go Go!!!")
		c <- true //向 channel 中存入一个值
	}()
	//当程序执行完毕之后再从通道中取出刚才赋的值
	<-c
	/**
	主线程启动了一个匿名子线程后就执行到了：<-c , 到达这里主线程就被阻塞了。只有当子线程向通道放入值后主线程阻塞才会被释放
	其实这个就是完成了消息的发送
	*/
}

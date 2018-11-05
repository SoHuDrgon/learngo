package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("当前系统核数：", runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU()) //设置当前程序执行使用的并发数
	//定义一个阻塞channel
	c := make(chan bool)
	//这里启动10个线程运行
	for i := 0; i < 10; i++ {
		go goRun(c, i)
	}
	//我们知道一共有10次循环，那么在这里就取10次,那么子线程goRun只有都执行完了主线程取才能完毕，因为这里也循环取10次，不够的话会被阻塞
	for i := 0; i < 10; i++ {
		<-c
	}
}

func goRun(c chan bool, index int) {
	a := 1
	//循环叠加1千万次并返回最终结果
	for i := 0; i < 10000000; i++ {
		a += i
	}
	fmt.Println("线程序号:", index, a) //往阻塞队列插入内容
	c <- true
}

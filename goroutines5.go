package main

import (
	"fmt"
)

/**
数据接收处理
*/
func main() {
	//批量初始化channel
	c1, c2 := make(chan int), make(chan string)
	//创建一个启动goroutine的匿名函数
	go func() {
		/**
		创建一个无限循环语句,使用select进行处理
		我们一般都是使用这种方式来处理不断的消息发送和处理
		*/
		for {
			select {
			case v, ok := <-c1:
				if !ok {
					break
				}
				fmt.Println("c1:", v)
			case v, ok := <-c2:
				if !ok {
					break
				}
				fmt.Println("c2:", v)
			}
		}
	}()

	c1 <- 1
	c2 <- "liang"
	c1 <- 2
	c2 <- "xuli"

	//关闭channel
	close(c1)
	close(c2)
}

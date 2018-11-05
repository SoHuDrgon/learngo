package main

import (
	"fmt"
)

/**
数据接收处理，这里实现随机接收0、1 数字并打印
*/
func main() {
	c := make(chan int)
	num := 0
	//创建一个启动goroutine的匿名函数
	go func() {
		for v := range c {
			num++
			if num&15 == 0 {
				fmt.Println()
			}
			fmt.Print(v, " ")
		}
	}()

	for {
		select {
		case c <- 0:
		case c <- 1:
		}
	}
}

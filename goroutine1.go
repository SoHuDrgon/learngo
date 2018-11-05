package main

import (
	"fmt"
	"time"
)

func main() {
	//启用一个goroutine
	go GoRun()
	//这里加一个休眠是因为主线程已启动就执行完毕消亡来，子线程还来不及执行
	time.Sleep(2 * time.Second)
}

func GoRun() {
	fmt.Println("Go Go Go!!!")
}

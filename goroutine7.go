package main

import (
	"fmt"
	"time"
)

/**
select的超时应用
*/
func main() {
	c := make(chan bool)
	select {
	case v := <-c:
		fmt.Println(v)
	case <-time.After(3 * time.Second):
		fmt.Println("TimeOut!!!")
	}
}

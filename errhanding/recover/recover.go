package main

import (
	"fmt"
)

func tryRecover() {
	defer func() {
		r := recover() //是任何类型
		if err, ok := r.(error); ok {
			fmt.Println("Error occurred:", err)
		} else {
			panic(fmt.Sprintf("I don't know what to do: %v", r))
		}
	}()
	//panic(errors.New("This is an error."))
	//a := 5 / 0
	//fmt.Println(a) //division by zero
	//b := 0
	//a := 5 / b
	//fmt.Println(a) //Error occurred: runtime error: integer divide by zero
	panic(123) //panic: I don't know what to do: 123

}
func main() {
	tryRecover()
}

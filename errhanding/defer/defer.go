package main

import (
	"bufio"
	"fmt"
	"github.com/guogang1990/learngo/functional/fib"
	"os"
)

func tryDefer() {
	//defer有个栈，是先进后出的，所以先执行2，后执行1
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	//defer不惧panic，return
	panic("error occurred.")
	return
}

func writeFile(filename string) {
	//file, err := os.Create(filename)
	//这里文件存在会报错，我们来处理错误
	file, err := os.OpenFile(
		filename, os.O_EXCL|os.O_CREATE, 0666)
	//自建error
	//err = errors.New("this is a custom error.")
	if err != nil {
		//panic(err)
		//fmt.Println("File already exists.")
		//fmt.Println("Error:", err.Error())
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Println(pathError.Op,
				pathError.Path,
				pathError.Err)
		}
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	f := fib.Fibnoacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}
func main() {
	writeFile("fib.txt")
	//tryDefer()
}

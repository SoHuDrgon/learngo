package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

//斐波那契数列实现
//1 1 2 3 5 8 13 21
//    a b
//      a b
//func fibnoacci() func() int{
func fibnoacci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

//定义接口
//先定义类型
type intGen func() int

func (g intGen) Read(
	p []byte) (n int, err error) {
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	//把数转成字符串，然后把s写进byte里面
	s := fmt.Sprintf("%d\n", next)
	//TODO: incorrect if p is too small!
	return strings.NewReader(s).Read(p)
}
func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
func main() {
	f := fibnoacci()
	//fmt.Println(f()) //1
	//fmt.Println(f()) //1
	//fmt.Println(f()) //2
	//fmt.Println(f()) //3
	//fmt.Println(f()) //5
	//fmt.Println(f()) //8
	//fmt.Println(f()) //13
	//fmt.Println(f()) //21
	printFileContents(f)

}

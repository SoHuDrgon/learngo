package basic

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//for循环语句
//for的条件里不需要括号
//for的条件里可以省略初始条件，结束条件，递增表达式

func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

//go语言没有while语句
func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

//死循环
//func forever() {
//	for {
//		fmt.Println("abc")
//	}
//}
func main() {
	fmt.Print(
		convertToBin(5),  //101
		convertToBin(13), //1011 --> 1101
	)
	printFile("abc.txt")
}

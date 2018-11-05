package _func

//返回值类型写在最后面
//可返回多个值
//函数作为参数
//没有默认参数，可选参数
import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		q, _ := div(a, b)
		return q
	default:
		panic("unsupported operation:" + op)
	}
}

func div(a, b int) (q, r int) {
	return a / b, a % b
}

//函数返回多个值时可以器名字
//仅用于非常简单的函数
//对于调用者而言没有区别
func div1(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}
func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args"+"(%d, %d)\n", opName, a, b)
	return op(a, b)
}

//可变参数列表
func sum(numbers ...int) int {
	sum := 0
	for i := range numbers {
		sum += numbers[i]
	}
	return sum
}

//func swap(a, b *int) {
//	*b, *a = *a, *b
//}
func swap(a, b int) (int, int) {
	return b, a
}
func main() {
	fmt.Println(eval(3, 4, "*"))
	fmt.Println(eval(3, 4, "/"))
	q, r := div1(13, 3)
	fmt.Println(q, r)
	fmt.Println(div(13, 3))
	fmt.Println(apply(pow, 3, 4))
	fmt.Println(apply(
		func(a int, b int) int {
			return int(math.Pow(float64(a), float64(b)))
		}, 3, 4))
	fmt.Println(sum(1, 2, 3, 4, 5, 6))
	a, b := 3, 4
	//swap(&a, &b)
	a, b = swap(a, b)
	fmt.Println(a, b)
}

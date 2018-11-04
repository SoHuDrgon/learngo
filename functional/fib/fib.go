package fib

//斐波那契数列实现
//1 1 2 3 5 8 13 21
//    a b
//      a b
func Fibnoacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

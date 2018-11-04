package basic

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "yes,我喜欢go语言！"
	fmt.Println(len(s)) //24
	fmt.Println(s)
	fmt.Printf("%s\n", []byte(s))
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()
	//79 65 73 2C E6 88 91 E5 96 9C E6 AC A2 67 6F E8 AF AD E8 A8 80 EF BC 81
	for i, ch := range s { //ch is a rune
		fmt.Printf("(%d %X)", i, ch)
	}
	fmt.Println()
	//(0 79)(1 65)(2 73)(3 2C)(4 6211)(7 559C)(10 6B22)(13 67)(14 6F)(15 8BED)(18 8A00)(21 FF01)
	fmt.Println("Rune count:",
		utf8.RuneCountInString(s))
	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c) ", i, ch)
	}
	fmt.Println()
}

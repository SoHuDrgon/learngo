package main

import "fmt"

var lastOccurred = make([]int, 0xffff)

func LengthOfNonRepeatingSubStr(s string) int {
	//map花费时间太多，我们来优化一下，将map换成数组两字节来测试
	//lastOccurred := make(map[rune]int)
	//stores last occurred pos + 1.
	// 0 means not seen.
	//每次都要去内存里面开辟空间，30000次后内存就频繁垃圾回收gc了
	//我们需要把它拿出去,拿出去后我们还需要清理
	//lastOccurred := make([]int, 0xffff)
	//这里会clear memory
	for i := range lastOccurred {
		lastOccurred[i] = 0
	}
	start := 0
	maxLength := 0
	for i, ch := range []rune(s) {
		//if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
		//start = lastI + 1
		if lastI := lastOccurred[ch]; lastI > start {
			start = lastI
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		//lastOccurred[ch] = i
		lastOccurred[ch] = i + 1
	}
	return maxLength
}
func main() {
	fmt.Println(LengthOfNonRepeatingSubStr("abcabcbb"))
	fmt.Println(LengthOfNonRepeatingSubStr("bbbbb"))
	fmt.Println(LengthOfNonRepeatingSubStr("pwwkew"))
	fmt.Println(LengthOfNonRepeatingSubStr(""))
	fmt.Println(LengthOfNonRepeatingSubStr("b"))
	fmt.Println(LengthOfNonRepeatingSubStr("fijklmn"))
	fmt.Println(LengthOfNonRepeatingSubStr("我爱go语言！"))
	fmt.Println(LengthOfNonRepeatingSubStr("我要学会go语言！"))
	fmt.Println(LengthOfNonRepeatingSubStr("一二三二一"))
}

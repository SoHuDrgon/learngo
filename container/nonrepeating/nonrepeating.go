package main

import "fmt"

func lengthOfNonRepeatingSubStr1(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0
	for i, ch := range []byte(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func main() {
	fmt.Println(lengthOfNonRepeatingSubStr1("abcabcbb"))
	fmt.Println(lengthOfNonRepeatingSubStr1("bbbbb"))
	fmt.Println(lengthOfNonRepeatingSubStr1("pwwkew"))
	fmt.Println(lengthOfNonRepeatingSubStr1(""))
	fmt.Println(lengthOfNonRepeatingSubStr1("b"))
	fmt.Println(lengthOfNonRepeatingSubStr1("fijklmn"))
}

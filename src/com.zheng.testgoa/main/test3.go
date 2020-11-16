package main

import (
	"fmt"
	"strconv"
)

func main() {

	var istrue = isPalindrome(121)
	fmt.Println(istrue)
}

func isPalindrome(x int) bool {
	var str string = strconv.Itoa(x)
	var length int = len(str)
	var intarray [length]int

	for i := 0; i < len(str); i++ {
		intarray[i] = x % exponent(10, i)
	}

	for i, j := 0, len(intarray)-1; i < j; i, j = i+1, j-1 {
		intarray[i], intarray[j] = intarray[j], intarray[i]
	}

	if str == resultstr {
		return true
	} else {
		return false
	}

}

func exponent(a, n int) int {
	result := int(1)
	for i := n; i > 0; i >>= 1 {
		if i&1 != 0 {
			result *= a
		}
		a *= a
	}
	return result
}

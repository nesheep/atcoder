package main

import (
	"fmt"
	"os"
)

func isPalindrome(s []byte) bool {
	for i := 0; i < len(s)/2; i++ {
		j := len(s) - 1 - i
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func main() {
	var s []byte
	fmt.Scan(&s)

	for i := len(s); i > 0; i-- {
		for j := 0; j <= len(s)-i; j++ {
			if isPalindrome(s[j : j+i]) {
				fmt.Println(i)
				os.Exit(0)
			}
		}
	}
}

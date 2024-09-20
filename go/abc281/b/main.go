package main

import (
	"fmt"
	"strconv"
)

func judge(s string) bool {
	if len(s) != 8 {
		return false
	}
	if s[0] < 'A' || s[0] > 'Z' {
		return false
	}
	if n, err := strconv.Atoi(s[1:7]); err != nil || n < 100000 {
		return false
	}
	if s[7] < 'A' || s[7] > 'Z' {
		return false
	}
	return true
}

func main() {
	var s string
	fmt.Scan(&s)

	if judge(s) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

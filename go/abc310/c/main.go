package main

import (
	"bufio"
	"fmt"
	"os"
)

func minString(s, t string) string {
	if s < t {
		return s
	}
	return t
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	rd := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(rd, &n)

	sm := make(map[string]bool)
	for i := 0; i < n; i++ {
		var s string
		fmt.Fscan(rd, &s)
		sm[minString(s, reverseString(s))] = true
	}

	fmt.Println(len(sm))
}

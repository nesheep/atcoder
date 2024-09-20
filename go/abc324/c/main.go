package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func judge(s, t []byte) bool {
	if len(s) == len(t) {
		diff := 0
		for i := 0; i < len(s); i++ {
			if s[i] != t[i] {
				diff++
			}
		}
		return diff <= 1
	} else if len(s) == len(t)+1 {
		for i := 0; i < len(t); i++ {
			if s[i] != t[i] {
				return string(s[i+1:]) == string(t[i:])
			}
		}
		return true
	} else if len(s)+1 == len(t) {
		for i := 0; i < len(s); i++ {
			if s[i] != t[i] {
				return string(s[i:]) == string(t[i+1:])
			}
		}
		return true
	}
	return false
}

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var n int
	var t []byte
	fmt.Fscan(rd, &n, &t)

	s := make([][]byte, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &s[i])
	}

	ans := make([]string, 0, n)
	for i := 0; i < n; i++ {
		if judge(t, s[i]) {
			ans = append(ans, strconv.Itoa(i+1))
		}
	}

	fmt.Fprintln(wr, len(ans))
	fmt.Fprintln(wr, strings.Join(ans, " "))
}

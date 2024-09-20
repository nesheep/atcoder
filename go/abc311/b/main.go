package main

import (
	"bufio"
	"fmt"
	"os"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	rd := bufio.NewReader(os.Stdin)

	var n, d int
	fmt.Fscan(rd, &n, &d)

	s := make([][]byte, n)
	for i := range s {
		fmt.Fscan(rd, &s[i])
	}

	a := make([]bool, d)
	for i := range a {
		a[i] = true
	}

	for i := range s {
		for j := range s[i] {
			if s[i][j] == 'x' {
				a[j] = false
			}
		}
	}

	var m, t int
	for i := range a {
		if a[i] {
			t++
		} else {
			t = 0
		}
		m = max(m, t)
	}

	fmt.Println(m)
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func Permute(s [][]byte, f func([][]byte)) {
	permute(s, 0, f)
}

func permute(s [][]byte, i int, f func([][]byte)) {
	if i > len(s) {
		f(s)
		return
	}
	permute(s, i+1, f)
	for j := i + 1; j < len(s); j++ {
		s[i], s[j] = s[j], s[i]
		permute(s, i+1, f)
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var n, m int
	fmt.Fscan(rd, &n, &m)

	s := make([][]byte, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &s[i])
	}

	ans := "No"
	Permute(s, func(s [][]byte) {
		f := true
		for i := 0; i < n-1; i++ {
			var diff int
			for j := 0; j < m; j++ {
				if s[i][j] != s[i+1][j] {
					diff++
				}
			}
			if diff != 1 {
				f = false
				break
			}
		}
		if f {
			ans = "Yes"
		}
	})

	fmt.Fprintln(wr, ans)
}

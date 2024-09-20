package main

import (
	"bufio"
	"fmt"
	"os"
)

func judge(s [][]byte, i, j int) bool {
	for k := 0; k < 4; k++ {
		for l := 0; l < 4; l++ {
			var want byte = '#'
			if k == 3 || l == 3 {
				want = '.'
			}
			if s[i+k][j+l] != want || s[i+8-k][j+8-l] != want {
				return false
			}
		}
	}
	return true
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

	for i := 0; i < n-8; i++ {
		for j := 0; j < m-8; j++ {
			if judge(s, i, j) {
				fmt.Fprintln(wr, i+1, j+1)
			}
		}
	}
}

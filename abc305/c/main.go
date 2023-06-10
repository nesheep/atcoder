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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var h, w int
	fmt.Fscan(rd, &h, &w)

	ss := make([][]byte, h)
	for i := 0; i < h; i++ {
		fmt.Fscan(rd, &ss[i])
	}

	a, c := 500, 500
	b, d := 0, 0

	for i := range ss {
		for j := range ss[i] {
			if ss[i][j] == '#' {
				a = min(a, i)
				b = max(b, i)
				c = min(c, j)
				d = max(d, j)
			}
		}
	}

	for i := a; i <= b; i++ {
		for j := c; j <= d; j++ {
			if ss[i][j] == '.' {
				fmt.Fprintln(wr, i+1, j+1)
				return
			}
		}
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var r, c int
	fmt.Fscan(rd, &r, &c)

	b := make([][]byte, r)
	for i := range b {
		fmt.Fscan(rd, &b[i])
	}

	for i := range b {
		for j := range b[i] {
			if b[i][j] >= '1' && b[i][j] <= '9' {
				p := int(b[i][j] - '0')
				b[i][j] = '.'
				for k := i - p; k <= i+p; k++ {
					for l := j - p; l <= j+p; l++ {
						if k >= 0 && k < r && l >= 0 && l < c {
							if abs(i-k)+abs(j-l) <= p && b[k][l] == '#' {
								b[k][l] = '.'
							}
						}
					}
				}
			}
		}
	}

	for i := range b {
		fmt.Fprintln(wr, string(b[i]))
	}
}

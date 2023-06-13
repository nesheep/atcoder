package main

import (
	"bufio"
	"fmt"
	"os"
)

type pair struct {
	a, b int
}

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var n, q int
	fmt.Fscan(rd, &n, &q)

	m := map[pair]bool{}
	var t, a, b int
	for i := 0; i < q; i++ {
		fmt.Fscan(rd, &t, &a, &b)
		switch t {
		case 1:
			m[pair{a, b}] = true
		case 2:
			m[pair{a, b}] = false
		default:
			if m[pair{a, b}] && m[pair{b, a}] {
				fmt.Fprintln(wr, "Yes")
			} else {
				fmt.Fprintln(wr, "No")
			}
		}
	}
}

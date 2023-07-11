package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var n, q int
	fmt.Fscan(rd, &n, &q)

	a := make([][]int, n)
	for i := range a {
		var l int
		fmt.Fscan(rd, &l)
		a[i] = make([]int, l)
		for j := range a[i] {
			fmt.Fscan(rd, &a[i][j])
		}
	}

	for i := 0; i < q; i++ {
		var s, t int
		fmt.Fscan(rd, &s, &t)
		fmt.Fprintln(wr, a[s-1][t-1])
	}
}

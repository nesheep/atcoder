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

	y := make([]bool, n)
	r := make([]bool, n)

	var c, x int
	for i := 0; i < q; i++ {
		fmt.Fscan(rd, &c, &x)
		switch c {
		case 1:
			if y[x-1] {
				r[x-1] = true
			} else {
				y[x-1] = true
			}
		case 2:
			r[x-1] = true
		default:
			if r[x-1] {
				fmt.Fprintln(wr, "Yes")
			} else {
				fmt.Fprintln(wr, "No")
			}
		}
	}
}

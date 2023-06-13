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

	var n int
	fmt.Fscan(rd, &n)

	as := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(rd, &as[i])
	}

	var q int
	fmt.Fscan(rd, &q)

	b := -1
	c := map[int]int{}
	var qq, iq, xq int
	for i := 0; i < q; i++ {
		fmt.Fscan(rd, &qq)
		switch qq {
		case 1:
			fmt.Fscan(rd, &xq)
			b = xq
			c = map[int]int{}
		case 2:
			fmt.Fscan(rd, &iq, &xq)
			c[iq] += xq
		default:
			fmt.Fscan(rd, &iq)
			if b < 0 {
				fmt.Fprintln(wr, as[iq]+c[iq])
			} else {
				fmt.Fprintln(wr, b+c[iq])
			}
		}
	}
}

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

	var c, k, x int
	for i := 0; i < q; i++ {
		fmt.Fscan(rd, &c, &k)
		if c == 1 {
			fmt.Fscan(rd, &x)
			as[k] = x
		} else {
			fmt.Println(as[k])
		}
	}
}

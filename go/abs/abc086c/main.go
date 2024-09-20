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

	var n int
	fmt.Fscan(rd, &n)

	ans := true
	var t, x, y int
	var ti, xi, yi int
	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &ti, &xi, &yi)
		td := ti - t
		dd := abs(xi-x) + abs(yi-y)
		if dd > td || dd%2 != td%2 {
			ans = false
			break
		}
		t, x, y = ti, xi, yi
	}

	if ans {
		fmt.Fprintln(wr, "Yes")
	} else {
		fmt.Fprintln(wr, "No")
	}
}

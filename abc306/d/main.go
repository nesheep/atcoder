package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var n int
	fmt.Fscan(rd, &n)

	x := make([]bool, n)
	y := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &x[i], &y[i])
	}

	d := make([]int, n)
	e := make([]int, n)
	f := make([]int, n)
	g := make([]int, n)
	if x[0] {
		f[0] = math.MinInt64
		g[0] = y[0]
	} else {
		f[0] = y[0]
		g[0] = math.MinInt64
	}

	for i := 1; i < n; i++ {
		d[i] = max(d[i-1], f[i-1])
		e[i] = max(e[i-1], g[i-1])
		if x[i] {
			f[i] = math.MinInt64
			g[i] = max(d[i-1], f[i-1]) + y[i]
		} else {
			f[i] = max(max(d[i-1], e[i-1]), max(f[i-1], g[i-1])) + y[i]
			g[i] = math.MinInt64
		}
	}

	fmt.Fprintln(wr, max(max(d[n-1], e[n-1]), max(f[n-1], g[n-1])))
}

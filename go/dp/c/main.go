package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n int
	fmt.Fscan(r, &n)

	a := make([]int, n)
	b := make([]int, n)
	c := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &a[i])
		fmt.Fscan(r, &b[i])
		fmt.Fscan(r, &c[i])
	}

	dpa := make([]int, n)
	dpb := make([]int, n)
	dpc := make([]int, n)

	dpa[0] = a[0]
	dpb[0] = b[0]
	dpc[0] = c[0]
	for i := 1; i < n; i++ {
		dpa[i] = max(dpb[i-1], dpc[i-1]) + a[i]
		dpb[i] = max(dpc[i-1], dpa[i-1]) + b[i]
		dpc[i] = max(dpa[i-1], dpb[i-1]) + c[i]
	}

	fmt.Fprintln(w, max(max(dpa[n-1], dpb[n-1]), dpc[n-1]))
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

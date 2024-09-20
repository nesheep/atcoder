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

	h := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &h[i])
	}

	dp := make([]int, n)
	dp[1] = abs(h[0] - h[1])
	for i := 2; i < n; i++ {
		d1 := dp[i-2] + abs(h[i-2]-h[i])
		d2 := dp[i-1] + abs(h[i-1]-h[i])
		dp[i] = min(d1, d2)
	}

	fmt.Fprintln(w, dp[n-1])
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

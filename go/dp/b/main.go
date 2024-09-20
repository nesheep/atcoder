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

	var n, k int
	fmt.Fscan(r, &n, &k)

	h := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &h[i])
	}

	dp := make([]int, n)
	for i := 1; i < n; i++ {
		dp[i] = dp[i-1] + abs(h[i-1]-h[i])
		for j := 2; j <= k && j <= i; j++ {
			if j <= i {
				dp[i] = min(dp[i], dp[i-j]+abs(h[i-j]-h[i]))
			}
		}
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

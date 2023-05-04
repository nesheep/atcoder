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

	a := make([]int, n+1)
	for i := 1; i < n; i++ {
		fmt.Fscan(r, &a[i])
	}

	dp := make([]int, n+1)
	dp[1] = a[1]
	for i := 2; i < n+1; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+a[i])
	}

	fmt.Fprintln(w, dp[n])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

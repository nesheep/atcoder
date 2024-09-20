package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	var n, w int
	fmt.Fscan(r, &n, &w)

	ws := make([]int, n)
	vs := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &ws[i], &vs[i])
	}

	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, w+1)
	}

	dp[0][0] = 0
	for j := 1; j < w+1; j++ {
		dp[0][j] = -1
	}

	for i := 1; i < n+1; i++ {
		for j := 0; j < w+1; j++ {
			dp[i][j] = dp[i-1][j]
			if j >= ws[i-1] && dp[i-1][j-ws[i-1]] >= 0 {
				dp[i][j] = max(dp[i][j], dp[i-1][j-ws[i-1]]+vs[i-1])
			}
		}
	}

	var ans int
	for j := 0; j < w+1; j++ {
		ans = max(ans, dp[n][j])
	}

	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

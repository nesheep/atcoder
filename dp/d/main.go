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

	ws := make([]int, n+1)
	vs := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(r, &ws[i])
		fmt.Fscan(r, &vs[i])
	}

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, w+1)
	}

	for j := 1; j <= w; j++ {
		dp[0][j] = -1
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= w; j++ {
			dp[i][j] = dp[i-1][j]
			if j >= ws[i] && dp[i-1][j-ws[i]] >= 0 {
				dp[i][j] = max(dp[i][j], dp[i-1][j-ws[i]]+vs[i])
			}
		}
	}

	var ans int
	for j := 0; j <= w; j++ {
		ans = max(ans, dp[n][j])
	}

	fmt.Println(ans)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

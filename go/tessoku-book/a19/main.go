package main

import "fmt"

func main() {
	var n, w int
	fmt.Scan(&n, &w)

	ws := make([]int, n+1)
	vs := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&ws[i], &vs[i])
	}

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, w+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	dp[0][0] = 0

	for i := 1; i <= n; i++ {
		for j := 0; j <= w; j++ {
			dp[i][j] = dp[i-1][j]
			if j-ws[i] >= 0 && dp[i-1][j-ws[i]] >= 0 {
				dp[i][j] = max(dp[i][j], dp[i-1][j-ws[i]]+vs[i])
			}
		}
	}

	var ans int
	for _, v := range dp[n] {
		if v > ans {
			ans = v
		}
	}

	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

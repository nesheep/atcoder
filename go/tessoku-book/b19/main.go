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
		dp[i] = make([]int, 100_001)
		for j := range dp[i] {
			dp[i][j] = w + 1
		}
	}

	dp[0][0] = 0

	for i := 1; i <= n; i++ {
		for j := range dp[i] {
			dp[i][j] = dp[i-1][j]
			if j-vs[i] >= 0 && dp[i-1][j-vs[i]] >= 0 {
				dp[i][j] = min(dp[i][j], dp[i-1][j-vs[i]]+ws[i])
			}
		}
	}

	var ans int
	for j, v := range dp[n] {
		if v <= w {
			ans = j
		}
	}

	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

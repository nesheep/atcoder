package main

import "fmt"

func main() {
	var n int
	var s []byte
	fmt.Scan(&n, &s)

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if i == j {
				dp[i][j] = 1
				continue
			}
			if j-i == 1 && s[i] == s[j] {
				dp[i][j] = 2
				continue
			}
			dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			if s[i] == s[j] {
				dp[i][j] = max(dp[i][j], dp[i+1][j-1]+2)
			} else {
				dp[i][j] = max(dp[i][j], dp[i+1][j-1])
			}
		}
	}

	fmt.Println(dp[0][n-1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

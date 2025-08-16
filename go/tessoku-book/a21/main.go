package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	ps := make([]int, n)
	as := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&ps[i], &as[i])
	}

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := n - 1; j >= i; j-- {
			if i > 0 {
				dp[i][j] = max(dp[i][j], dp[i-1][j])
				if i <= ps[i-1]-1 && ps[i-1]-1 <= j {
					dp[i][j] = max(dp[i][j], dp[i-1][j]+as[i-1])
				}
			}
			if j < n-1 {
				dp[i][j] = max(dp[i][j], dp[i][j+1])
				if i <= ps[j+1]-1 && ps[j+1]-1 <= j {
					dp[i][j] = max(dp[i][j], dp[i][j+1]+as[j+1])
				}
			}
		}
	}

	var ans int
	for i := range dp {
		for j := range dp[i] {
			ans = max(ans, dp[i][j])
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

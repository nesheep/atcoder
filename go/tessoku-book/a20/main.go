package main

import "fmt"

func main() {
	var s, t []byte
	fmt.Scan(&s, &t)

	dp := make([][]int, len(s)+1)
	for i := range dp {
		dp[i] = make([]int, len(t)+1)
	}

	for i := range dp {
		for j := range dp[i] {
			if i == 0 || j == 0 {
				continue
			}
			dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			if s[i-1] == t[j-1] {
				dp[i][j] = max(dp[i][j], dp[i-1][j-1]+1)
			}
		}
	}

	fmt.Println(dp[len(s)][len(t)])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

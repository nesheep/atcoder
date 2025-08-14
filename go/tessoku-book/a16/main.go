package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	as := make([]int, n)
	for i := 1; i < n; i++ {
		fmt.Scan(&as[i])
	}

	bs := make([]int, n)
	for i := 2; i < n; i++ {
		fmt.Scan(&bs[i])
	}

	dp := make([]int, n)
	dp[0], dp[1] = 0, as[1]

	for i := 2; i < n; i++ {
		dp[i] = min(dp[i-1]+as[i], dp[i-2]+bs[i])
	}

	fmt.Println(dp[n-1])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

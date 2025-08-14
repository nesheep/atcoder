package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	hs := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&hs[i])
	}

	dp := make([]int, n)
	dp[1] = cost(hs, 0, 1)

	for i := 2; i < n; i++ {
		c1 := dp[i-1] + cost(hs, i-1, i)
		c2 := dp[i-2] + cost(hs, i-2, i)
		dp[i] = min(c1, c2)
	}

	fmt.Println(dp[n-1])
}

func cost(hs []int, i, j int) int {
	c := hs[i] - hs[j]
	if c < 0 {
		return -c
	}
	return c
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

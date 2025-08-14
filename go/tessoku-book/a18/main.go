package main

import "fmt"

func main() {
	var n, s int
	fmt.Scan(&n, &s)

	as := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&as[i])
	}

	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, s+1)
	}

	dp[0][0] = true
	if as[0] <= s {
		dp[0][as[0]] = true
	}

	for i := 1; i < n; i++ {
		for j := 0; j <= s; j++ {
			if !dp[i-1][j] {
				continue
			}
			dp[i][j] = true
			if j+as[i] <= s {
				dp[i][j+as[i]] = true
			}
		}
	}

	var ans bool
	for i := 0; i < n; i++ {
		if dp[i][s] {
			ans = true
			break
		}
	}

	if ans {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

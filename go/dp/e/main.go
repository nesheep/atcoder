package main

import (
	"bufio"
	"fmt"
	"os"
)

const inf = 1000000000000

func main() {
	r := bufio.NewReader(os.Stdin)

	var n, w int
	fmt.Fscan(r, &n, &w)

	ws := make([]int, n+1)
	vs := make([]int, n+1)
	var vsum int
	for i := 1; i <= n; i++ {
		fmt.Fscan(r, &ws[i])
		fmt.Fscan(r, &vs[i])
		vsum += vs[i]
	}

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, vsum+1)
		for j := 1; j <= vsum; j++ {
			dp[i][j] = inf
		}
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= vsum; j++ {
			dp[i][j] = dp[i-1][j]
			if j >= vs[i] {
				dp[i][j] = min(dp[i][j], dp[i-1][j-vs[i]]+ws[i])
			}
		}
	}

	var ans int
	for j := 0; j <= vsum; j++ {
		if dp[n][j] <= w {
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

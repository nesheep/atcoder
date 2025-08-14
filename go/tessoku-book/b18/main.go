package main

import (
	"fmt"
	"strconv"
	"strings"
)

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

	var exists bool
	for i := 0; i < n; i++ {
		if dp[i][s] {
			exists = true
			break
		}
	}

	if !exists {
		fmt.Println(-1)
		return
	}

	var p []int
	i, j := n-1, s
	for {
		if i == 0 {
			if j != 0 {
				p = append(p, i)
			}
			break
		}
		if !dp[i-1][j] {
			p = append(p, i)
			j = j - as[i]
		}
		i--
	}

	ans := make([]string, 0, len(p))
	for i := len(p) - 1; i >= 0; i-- {
		ans = append(ans, strconv.Itoa(p[i]+1))
	}

	fmt.Println(len(ans))
	fmt.Println(strings.Join(ans, " "))
}

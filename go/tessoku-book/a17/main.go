package main

import (
	"fmt"
	"strconv"
	"strings"
)

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

	p := make([]int, 0, n)
	for i := n - 1; ; {
		p = append(p, i)
		if i == 0 {
			break
		}
		if dp[i]-as[i] == dp[i-1] {
			i--
		} else {
			i = i - 2
		}
	}

	ans := make([]string, 0, len(p))
	for i := len(p) - 1; i >= 0; i-- {
		ans = append(ans, strconv.Itoa(p[i]+1))
	}

	fmt.Println(len(ans))
	fmt.Println(strings.Join(ans, " "))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

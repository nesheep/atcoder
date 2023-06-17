package main

import (
	"bufio"
	"fmt"
	"os"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var n int
	fmt.Fscan(rd, &n)

	var x bool
	var y int
	dp := [2]int{}
	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &x, &y)
		if x {
			dp[1] = max(dp[1], dp[0]+y)
		} else {
			dp[0] = max(max(dp[0], dp[0]+y), dp[1]+y)
		}
	}

	fmt.Fprintln(wr, max(dp[0], dp[1]))
}

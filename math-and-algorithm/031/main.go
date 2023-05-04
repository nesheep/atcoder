package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n int
	fmt.Fscan(r, &n)

	a := make([]int, n+1)
	for i := 1; i < n; i++ {
		fmt.Fscan(r, &a[i])
	}

	b := true
	dp := make([]int, n+1)
	dp[1] = a[1]
	for i := 2; i < n+1; i++ {
		if !b {
			b = true
			dp[i] = dp[i-1] + a[i]
		} else {
			if dp[i-2]+a[i] < dp[i-1] {
				b = false
				dp[i] = dp[i-1]
			} else {
				dp[i] = dp[i-2] + a[i]
			}
		}
	}

	fmt.Fprintln(w, dp[n])
}

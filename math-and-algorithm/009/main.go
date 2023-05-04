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

	var n, s int
	fmt.Fscan(r, &n, &s)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &a[i])
	}

	dp := make([][]bool, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]bool, s+1)
	}

	dp[0][0] = true
	for i := 1; i < n+1; i++ {
		for j := 0; j < s+1; j++ {
			dp[i][j] = dp[i-1][j] || j-a[i-1] >= 0 && dp[i-1][j-a[i-1]]
		}
	}

	if dp[n][s] {
		fmt.Fprintln(w, "Yes")
	} else {
		fmt.Fprintln(w, "No")
	}
}

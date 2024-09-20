package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var n, k int
	fmt.Fscan(rd, &n, &k)

	as := make([]string, 0, n)
	var a string
	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &a)
		if i >= k {
			as = append(as, a)
		}
	}
	for i := 0; i < min(n, k); i++ {
		as = append(as, "0")
	}

	fmt.Fprintln(wr, strings.Join(as, " "))
}

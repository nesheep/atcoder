package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rd := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(rd, &n, &m)

	d := make(map[int]bool, m)
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Fscan(rd, &a, &b)
		d[b] = true
	}

	if len(d) < n-1 {
		fmt.Println(-1)
		os.Exit(0)
	}

	for i := 1; i <= n; i++ {
		if !d[i] {
			fmt.Println(i)
			os.Exit(0)
		}
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var n, m int
	fmt.Fscan(rd, &n, &m)

	g := make([][]int, n)
	var a, b int
	for i := 0; i < m; i++ {
		fmt.Fscan(rd, &a, &b)
		g[a-1] = append(g[a-1], b-1)
		g[b-1] = append(g[b-1], a-1)
	}

	var ans int
	for i := range g {
		var cnt int
		for j := range g[i] {
			if g[i][j] < i {
				cnt++
			}
		}
		if cnt == 1 {
			ans++
		}
	}

	fmt.Fprintln(wr, ans)
}

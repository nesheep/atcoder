package main

import (
	"bufio"
	"fmt"
	"os"
)

func dfs(g [][]int, v int, color []int) bool {
	for _, u := range g[v] {
		if color[u] == 0 {
			color[u] = 3 - color[v]
			if !dfs(g, u, color) {
				return false
			}
		} else {
			if color[v] == color[u] {
				return false
			}
		}
	}
	return true
}

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

	color := make([]int, n)
	for i := range color {
		if color[i] == 0 {
			color[i] = 1
			if !dfs(g, i, color) {
				fmt.Fprintln(wr, "No")
				return
			}
		}
	}

	fmt.Fprintln(wr, "Yes")
}

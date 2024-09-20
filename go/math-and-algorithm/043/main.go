package main

import (
	"bufio"
	"fmt"
	"os"
)

func connected(g [][]int) bool {
	visited := make([]bool, len(g))

	var dfs func(p int)
	dfs = func(p int) {
		visited[p] = true
		for _, q := range g[p] {
			if !visited[q] {
				dfs(q)
			}
		}
	}

	dfs(1)

	for _, v := range visited[1:] {
		if !v {
			return false
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

	g := make([][]int, n+1)
	var a, b int
	for i := 0; i < m; i++ {
		fmt.Fscan(rd, &a, &b)
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}

	if connected(g) {
		fmt.Fprintln(wr, "The graph is connected.")
	} else {
		fmt.Fprintln(wr, "The graph is not connected.")
	}
}

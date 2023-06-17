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

	var n int
	fmt.Fscan(rd, &n)

	g := make(map[int][]int)
	var a, b int
	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &a, &b)
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}

	visited := make(map[int]bool)

	var dfs func(int)
	dfs = func(v int) {
		visited[v] = true
		for _, u := range g[v] {
			if !visited[u] {
				dfs(u)
			}
		}
	}

	dfs(1)

	var ans int
	for k := range visited {
		if k > ans {
			ans = k
		}
	}

	fmt.Fprintln(wr, ans)
}

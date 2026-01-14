package main

import (
	"fmt"
	"slices"
)

func main() {
	_, _, g := input()
	ans := isConnected(g)
	output(ans)
}

func isConnected(g [][]int) bool {
	acc := make([]bool, len(g))
	dfs(g, 0, acc)
	return !slices.Contains(acc, false)
}

func dfs(g [][]int, i int, acc []bool) {
	acc[i] = true
	for _, v := range g[i] {
		if !acc[v] {
			dfs(g, v, acc)
		}
	}
}

func input() (int, int, [][]int) {
	var n, m int
	fmt.Scan(&n, &m)
	g := make([][]int, n)
	for range m {
		var a, b int
		fmt.Scan(&a, &b)
		g[a-1] = append(g[a-1], b-1)
		g[b-1] = append(g[b-1], a-1)
	}
	return n, m, g
}

func output(ans bool) {
	if ans {
		fmt.Println("The graph is connected.")
	} else {
		fmt.Println("The graph is not connected.")
	}
}

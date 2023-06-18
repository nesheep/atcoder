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

	var n, d int
	fmt.Fscan(rd, &n, &d)

	x := make([]int, n)
	y := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &x[i], &y[i])
	}

	visited := make([]bool, n)

	var dfs func(int)
	dfs = func(v int) {
		visited[v] = true
		for u := range visited {
			if !visited[u] && (x[v]-x[u])*(x[v]-x[u])+(y[v]-y[u])*(y[v]-y[u]) <= d*d {
				dfs(u)
			}
		}
	}

	dfs(0)

	for _, v := range visited {
		if v {
			fmt.Fprintln(wr, "Yes")
		} else {
			fmt.Fprintln(wr, "No")
		}
	}
}

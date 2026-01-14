package main

import "fmt"

func main() {
	_, _, g := input()
	ans := bfs(g)
	output(ans)
}

func bfs(g [][]int) []int {
	ans := make([]int, len(g))
	for i := range ans {
		ans[i] = -1
	}
	ans[0] = 0

	for q := []int{0}; len(q) > 0; {
		v := q[0]
		q = q[1:]
		for _, u := range g[v] {
			if ans[u] == -1 {
				ans[u] = ans[v] + 1
				q = append(q, u)
			}
		}
	}
	return ans
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

func output(ans []int) {
	for _, v := range ans {
		fmt.Println(v)
	}
}

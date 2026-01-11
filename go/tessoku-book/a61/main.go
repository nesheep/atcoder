package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	_, _, g := input()
	output(g)
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

func output(g [][]int) {
	for i := range g {
		s := make([]string, 0, len(g[i]))
		for _, j := range g[i] {
			s = append(s, strconv.Itoa(j+1))
		}
		fmt.Printf("%d: {%s}\n", i+1, strings.Join(s, ", "))
	}
}

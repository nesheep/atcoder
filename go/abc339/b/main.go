package main

import (
	"bytes"
	"fmt"
)

func main() {
	var h, w, n int
	fmt.Scan(&h, &w, &n)

	g := make([][]byte, h)
	for i := range g {
		g[i] = bytes.Repeat([]byte("."), w)
	}

	ds := [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	p := [2]int{0, 0}
	d := 0
	for i := 0; i < n; i++ {
		if g[p[0]][p[1]] == '.' {
			g[p[0]][p[1]] = '#'
			d = (d + 1) % 4
		} else {
			g[p[0]][p[1]] = '.'
			d = (d - 1 + 4) % 4
		}
		p[0] = (p[0] + ds[d][0] + h) % h
		p[1] = (p[1] + ds[d][1] + w) % w
	}

	for i := range g {
		fmt.Printf("%s\n", g[i])
	}
}

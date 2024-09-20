package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rd := bufio.NewReader(os.Stdin)

	var h, w int
	fmt.Fscan(rd, &h, &w)

	sTmp := make([][]byte, h)
	for i := 0; i < h; i++ {
		fmt.Fscan(rd, &sTmp[i])
	}

	s := make([]byte, h*w)
	g := make([][]int, h*w)
	for r := 0; r < h; r++ {
		for c := 0; c < w; c++ {
			s[r*w+c] = sTmp[r][c]
			if r > 0 {
				g[r*w+c] = append(g[r*w+c], (r-1)*w+c)
			}
			if r < h-1 {
				g[r*w+c] = append(g[r*w+c], (r+1)*w+c)
			}
			if c > 0 {
				g[r*w+c] = append(g[r*w+c], r*w+c-1)
			}
			if c < w-1 {
				g[r*w+c] = append(g[r*w+c], r*w+c+1)
			}
		}
	}

	snuke := []byte{'s', 'n', 'u', 'k', 'e'}

	visited := make([]bool, h*w)

	var dfs func(p, q int)
	dfs = func(p, q int) {
		visited[p] = true
		for _, v := range g[p] {
			if !visited[v] && s[v] == snuke[q] {
				dfs(v, (q+1)%5)
			}
		}
	}

	if s[0] == snuke[0] {
		dfs(0, 1)
	}

	ans := "No"
	if visited[h*w-1] {
		ans = "Yes"
	}

	fmt.Println(ans)
}

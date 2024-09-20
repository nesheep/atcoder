package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var n, m int
	fmt.Fscan(rd, &n, &m)

	g := make([][]int, n)
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Fscan(rd, &a, &b)
		g[a-1] = append(g[a-1], b-1)
		g[b-1] = append(g[b-1], a-1)
	}

	for i := range g {
		sort.Ints(g[i])
		s := make([]string, 0, len(g[i]))
		for j := range g[i] {
			s = append(s, strconv.Itoa(g[i][j]+1))
		}
		ans := strconv.Itoa(len(g[i]))
		if len(g[i]) != 0 {
			ans = fmt.Sprintf("%s %s", ans, strings.Join(s, " "))
		}
		fmt.Fprintln(wr, ans)
	}
}

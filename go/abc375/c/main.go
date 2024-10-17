package main

import (
	"bufio"
	"fmt"
	"os"
)

type point [2]int

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func operationCnt(p point, n int) int {
	i, j := p[0], p[1]
	return min(min(i, n-1-i), min(j, n-1-j)) + 1
}

func operate(p point, n, cnt int) point {
	i, j := p[0], p[1]
	switch cnt % 4 {
	case 0:
		return point{i, j}
	case 1:
		return point{n - 1 - j, i}
	case 2:
		return point{n - 1 - i, n - 1 - j}
	case 3:
		return point{j, n - 1 - i}
	default:
		return point{0, 0}
	}
}

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var n int
	fmt.Fscan(rd, &n)

	a := make([][]byte, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &a[i])
	}

	ans := make([][]byte, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			p := point{i, j}
			cnt := operationCnt(p, n)
			q := operate(p, n, cnt)
			ans[i][j] = a[q[0]][q[1]]
		}
	}

	for _, v := range ans {
		fmt.Fprintln(wr, string(v))
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

type pair struct {
	x, y int
}

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var n, m, h, k int
	var s []byte
	fmt.Fscan(rd, &n, &m, &h, &k, &s)

	ms := make(map[pair]bool, n)
	var mx, my int
	for i := 0; i < m; i++ {
		fmt.Fscan(rd, &mx, &my)
		ms[pair{mx, my}] = true
	}

	t := pair{}
	for i := 0; i < n; i++ {
		h--
		switch s[i] {
		case 'R':
			t.x++
		case 'L':
			t.x--
		case 'U':
			t.y++
		case 'D':
			t.y--
		}
		if h < 0 {
			fmt.Fprintln(wr, "No")
			return
		}
		if ms[t] && h < k {
			ms[t] = false
			h = k
		}
	}

	fmt.Fprintln(wr, "Yes")
}

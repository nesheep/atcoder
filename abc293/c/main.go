package main

import (
	"bufio"
	"fmt"
	"os"
)

func Combination(n, k int, f func([]int)) {
	ptn := make([]int, k)

	var c func(p, s int)
	c = func(p, s int) {
		if p == k {
			f(ptn)
			return
		}
		for i := s; i < n+p-k+1; i++ {
			ptn[p] = i
			c(p+1, i+1)
		}
	}

	c(0, 0)
}

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var h, w int
	fmt.Fscan(rd, &h, &w)

	s := make([][]int, h)
	for i := range s {
		s[i] = make([]int, w)
		for j := range s[i] {
			fmt.Fscan(rd, &s[i][j])
		}
	}

	var ans int
	Combination(h+w-2, h-1, func(ptn []int) {
		ptnM := make(map[int]bool, len(ptn))
		for _, p := range ptn {
			ptnM[p] = true
		}
		m := make(map[int]bool, h+w-2)
		var x, y int
		for i := 0; i < h+w-2; i++ {
			m[s[x][y]] = true
			if ptnM[i] {
				x++
			} else {
				y++
			}
			if m[s[x][y]] {
				return
			}
		}
		ans++
	})

	fmt.Fprintln(wr, ans)
}

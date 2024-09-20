package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type pair struct {
	x, y int
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var h, w int
	fmt.Fscan(rd, &h, &w)

	c := make(map[pair]bool, h*w)
	var t []byte
	for i := 0; i < h; i++ {
		fmt.Fscan(rd, &t)
		for j := 0; j < w; j++ {
			if t[j] == '#' {
				c[pair{i, j}] = true
			}
		}
	}

	n := min(h, w)
	s := make([]int, n+1)
	for k := range c {
		var cnt int
		for i := 1; i <= n/2; i++ {
			c1 := c[pair{k.x + i, k.y + i}]
			c2 := c[pair{k.x - i, k.y + i}]
			c3 := c[pair{k.x + i, k.y - i}]
			c4 := c[pair{k.x - i, k.y - i}]
			if !c1 || !c2 || !c3 || !c4 {
				break
			}
			cnt++
		}
		s[cnt]++
	}

	ss := make([]string, 0, n)
	for i := 1; i <= n; i++ {
		ss = append(ss, strconv.Itoa(s[i]))
	}

	fmt.Fprintln(wr, strings.Join(ss, " "))
}

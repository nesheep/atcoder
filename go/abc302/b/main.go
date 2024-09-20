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

	var h, w int
	fmt.Fscan(rd, &h, &w)

	ss := make([][]byte, h)
	for i := 0; i < h; i++ {
		fmt.Fscan(rd, &ss[i])
	}

	snuke := [5]byte{'s', 'n', 'u', 'k', 'e'}
	dx := [8]int{1, 1, 1, 0, 0, -1, -1, -1}
	dy := [8]int{1, 0, -1, 1, -1, 1, 0, -1}

	for x := 0; x < h; x++ {
		for y := 0; y < w; y++ {
			for d := 0; d < 8; d++ {
				ex := x + dx[d]*4
				ey := y + dy[d]*4
				if ex < 0 || ex >= h || ey < 0 || ey >= w {
					continue
				}
				ans := true
				for i, v := range snuke {
					if ss[x+dx[d]*i][y+dy[d]*i] != v {
						ans = false
						break
					}
				}
				if ans {
					for i := range snuke {
						fmt.Fprintln(wr, x+dx[d]*i+1, y+dy[d]*i+1)
					}
					return
				}
			}
		}
	}
}

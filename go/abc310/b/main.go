package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rd := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(rd, &n, &m)

	p := make([]int, n)
	c := make([]int, n)
	f := make([]map[int]bool, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &p[i], &c[i])
		f[i] = make(map[int]bool, c[i])
		for j := 0; j < c[i]; j++ {
			var tmp int
			fmt.Fscan(rd, &tmp)
			f[i][tmp] = true
		}
	}

	ans := "No"
All:
	for i := 0; i < n; i++ {
	One:
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			for k := range f[i] {
				if !f[j][k] {
					continue One
				}
			}
			if p[i] > p[j] || (p[i] == p[j] && len(f[i]) < len(f[j])) {
				ans = "Yes"
				break All
			}
		}
	}

	fmt.Println(ans)
}

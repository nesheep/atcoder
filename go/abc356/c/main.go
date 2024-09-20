package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rd := bufio.NewReader(os.Stdin)

	var n, m, k int
	fmt.Fscan(rd, &n, &m, &k)

	p := make([]int, 0, 1<<n)
	q := make([]bool, 0, 1<<n)

	for b := 0; b < 1<<n; b++ {
		p = append(p, b)
		q = append(q, true)
	}

	for i := 0; i < m; i++ {
		var c int
		fmt.Fscan(rd, &c)

		a := make([]int, c)
		for j := 0; j < c; j++ {
			fmt.Fscan(rd, &a[j])
		}

		var r string
		fmt.Fscan(rd, &r)

		for j := range p {
			if !q[j] {
				continue
			}

			if r == "o" {
				var cnt int
				for l := range a {
					if p[j]&(1<<(a[l]-1)) != 0 {
						cnt++
					}
				}

				if cnt < k {
					q[j] = false
				}
			} else {
				d := c - k + 1
				if d < 0 {
					continue
				}

				var dCnt int
				for l := range a {
					if p[j]&(1<<(a[l]-1)) == 0 {
						dCnt++
					}
				}

				if dCnt < d {
					q[j] = false
				}
			}
		}
	}

	var ans int
	for j := range q {
		if q[j] {
			ans++
		}
	}

	fmt.Println(ans)
}

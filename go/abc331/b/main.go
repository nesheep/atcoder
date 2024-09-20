package main

import "fmt"

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	var n, s, m, l int
	fmt.Scan(&n, &s, &m, &l)

	ans := 1 << 60
	for si := 0; si <= n; si++ {
		for mi := 0; mi <= n; mi++ {
			for li := 0; li <= n; li++ {
				if si*6+mi*8+li*12 >= n {
					ans = Min(ans, si*s+mi*m+li*l)
				}
			}
		}
	}

	fmt.Println(ans)
}

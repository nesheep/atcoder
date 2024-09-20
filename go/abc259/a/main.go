package main

import "fmt"

func main() {
	var n, m, x, t, d int
	fmt.Scan(&n, &m, &x, &t, &d)

	ans := t
	da := x - m
	if da > 0 {
		ans -= da * d
	}

	fmt.Println(ans)
}

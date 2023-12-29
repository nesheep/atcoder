package main

import "fmt"

func main() {
	var a, m, l, r int
	fmt.Scan(&a, &m, &l, &r)

	l -= a
	r -= a

	if l < 0 {
		x := -l/m + 1
		l += x * m
		r += x * m
	}

	fmt.Println(r/m - (l-1)/m)
}

package main

import "fmt"

type Matrix struct {
	a, b, c, d int
}

func (m Matrix) mul(n Matrix) Matrix {
	mod := 1000000000
	return Matrix{
		a: (m.a*n.a + m.b*n.c) % mod,
		b: (m.a*n.b + m.b*n.d) % mod,
		c: (m.c*n.a + m.d*n.c) % mod,
		d: (m.c*n.b + m.d*n.d) % mod,
	}
}

func (m Matrix) pow(n int) Matrix {
	p := m
	q := Matrix{1, 0, 0, 1}
	for i := 0; i < 60; i++ {
		if n&(1<<i) != 0 {
			q = q.mul(p)
		}
		p = p.mul(p)
	}
	return q
}

func main() {
	var n int
	fmt.Scan(&n)

	mod := 1000000000

	a := Matrix{1, 1, 1, 0}
	b := a.pow(n - 1)

	fmt.Println((b.c + b.d) % mod)
}

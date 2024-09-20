package main

import "fmt"

type Matrix struct {
	a, b, c, d int
}

func (m Matrix) Mul(n Matrix, mod int) Matrix {
	return Matrix{
		a: (m.a*n.a + m.b*n.c) % mod,
		b: (m.a*n.b + m.b*n.d) % mod,
		c: (m.c*n.a + m.d*n.c) % mod,
		d: (m.c*n.b + m.d*n.d) % mod,
	}
}

func (m Matrix) Pow(n int, mod int) Matrix {
	p := m
	q := Matrix{1, 0, 0, 1}
	for i := 0; i < 60; i++ {
		if n&(1<<i) != 0 {
			q = q.Mul(p, mod)
		}
		p = p.Mul(p, mod)
	}
	return q
}

func main() {
	var n int
	fmt.Scan(&n)

	mod := 1000000007

	a := Matrix{2, 1, 1, 0}
	b := a.Pow(n-1, mod)

	fmt.Println((b.c + b.d) % mod)
}

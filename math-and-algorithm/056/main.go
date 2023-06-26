package main

import "fmt"

type Matrix [3][3]int

func (m Matrix) Mul(n Matrix, mod int) Matrix {
	var o Matrix
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			o[i][j] = (m[i][0]*n[0][j] + m[i][1]*n[1][j] + m[i][2]*n[2][j]) % mod
		}
	}
	return o
}

func (m Matrix) Pow(n int, mod int) Matrix {
	p := m
	q := Matrix{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
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

	a := Matrix{{1, 1, 1}, {1, 0, 0}, {0, 1, 0}}
	b := a.Pow(n-1, mod)

	fmt.Println((2*b[2][0] + b[2][1] + b[2][2]) % mod)
}

package main

import "fmt"

func calc(a, b, c, x int) int {
	t := a + c
	q := x / t
	r := x % t
	if r > a {
		r = a
	}
	return (a*q + r) * b
}

func main() {
	var a, b, c, d, e, f, x int
	fmt.Scan(&a, &b, &c, &d, &e, &f, &x)

	td := calc(a, b, c, x)
	ad := calc(d, e, f, x)

	if td > ad {
		fmt.Println("Takahashi")
	} else if td < ad {
		fmt.Println("Aoki")
	} else {
		fmt.Println("Draw")
	}
}

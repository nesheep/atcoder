package main

import "fmt"

func gcd(a, b int) int {
	for a >= 1 && b >= 1 {
		if a < b {
			b %= a
		} else {
			a %= b
		}
	}
	if a >= 1 {
		return a
	}
	return b
}

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	res := a[0]
	for i := 1; i < n; i++ {
		res = gcd(res, a[i])
	}

	fmt.Println(res)
}

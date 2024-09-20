package main

import "fmt"

func modpow(a, b, mod int) int {
	p := a
	ans := 1
	for i := 0; i < 30; i++ {
		if b&(1<<i) != 0 {
			ans = (ans * p) % mod
		}
		p = (p * p) % mod
	}
	return ans
}

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	mod := 1000000007
	a, b := 1, 1
	for i := 1; i <= x+y; i++ {
		a = (a * i) % mod
	}
	for i := 1; i <= x; i++ {
		b = (b * i) % mod
	}
	for i := 1; i <= y; i++ {
		b = (b * i) % mod
	}

	fmt.Println((a * modpow(b, mod-2, mod)) % mod)
}

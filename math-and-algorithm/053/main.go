package main

import "fmt"

func modpow(a, b, mod int) int {
	p := a
	ans := 1
	for i := 0; i < 60; i++ {
		if b&(1<<i) != 0 {
			ans = (ans * p) % mod
		}
		p = (p * p) % mod
	}
	return ans
}

func main() {
	var n int
	fmt.Scan(&n)

	mod := 1000000007
	v := modpow(4, n+1, mod) - 1
	fmt.Println((v * modpow(3, mod-2, mod)) % mod)
}

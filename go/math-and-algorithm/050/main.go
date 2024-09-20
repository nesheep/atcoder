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
	var a, b int
	fmt.Scan(&a, &b)

	mod := 1000000007
	fmt.Println(modpow(a, b, mod))
}

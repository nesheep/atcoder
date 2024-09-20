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

	if 2*y-x < 0 || 2*x-y < 0 {
		fmt.Println(0)
		return
	}

	if (2*y-x)%3 != 0 || (2*x-y)%3 != 0 {
		fmt.Println(0)
		return
	}

	a, b := (2*y-x)/3, (2*x-y)/3
	mod := 1000000007
	c, d := 1, 1
	for i := 1; i <= a+b; i++ {
		c = (c * i) % mod
	}
	for i := 1; i <= a; i++ {
		d = (d * i) % mod
	}
	for i := 1; i <= b; i++ {
		d = (d * i) % mod
	}

	fmt.Println((c * modpow(d, mod-2, mod)) % mod)
}

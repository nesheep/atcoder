package main

import "fmt"

func mod(n, m int) int {
	if n < 0 {
		return n + m
	}
	return n % m
}

func main() {
	var a, b, c, d, e, f int
	fmt.Scan(&a, &b, &c, &d, &e, &f)

	m := 998244353
	fmt.Println(mod((((a%m)*(b%m))%m*(c%m))%m-(((d%m)*(e%m))%m*(f%m))%m, m))
}

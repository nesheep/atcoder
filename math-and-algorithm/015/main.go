package main

import "fmt"

func gcd(a, b int) int {
	var f func(a, b int) int
	f = func(a, b int) int {
		if b == 0 {
			return a
		}
		return f(b, a%b)
	}
	if a >= b {
		return f(a, b)
	}
	return f(b, a)
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(gcd(a, b))
}

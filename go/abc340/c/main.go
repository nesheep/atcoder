package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(F(n))
}

func F(n int) int {
	m := make(map[int]int)
	return f(n, m)
}

func f(n int, m map[int]int) int {
	if n == 1 {
		return 0
	}
	if m[n] == 0 {
		m[n] = n + f((n+1)/2, m) + f(n/2, m)
	}
	return m[n]
}

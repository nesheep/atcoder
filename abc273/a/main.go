package main

import "fmt"

func f(k int) int {
	if k == 0 {
		return 1
	}
	return k * f(k-1)
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(f(n))
}

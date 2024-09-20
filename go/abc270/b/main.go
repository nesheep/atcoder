package main

import "fmt"

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	var x, y, z int
	fmt.Scan(&x, &y, &z)

	if x*y < 0 || abs(x) < abs(y) {
		fmt.Println(abs(x))
	} else if y*z < 0 || abs(z) < abs(y) {
		fmt.Println(abs(z) + abs(x-z))
	} else {
		fmt.Println(-1)
	}
}

package main

import "fmt"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var n, x, y int
	fmt.Scan(&n, &x, &y)

	c1 := abs(x)+abs(y) <= n
	c2 := (x+y)%2 == n%2

	if c1 && c2 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

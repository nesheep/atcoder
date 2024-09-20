package main

import "fmt"

func pow(x, y int) int {
	if y == 0 {
		return 1
	}
	return x * pow(x, y-1)
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(pow(a, b) + pow(b, a))
}

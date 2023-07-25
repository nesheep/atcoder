package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var l1, r1, l2, r2 int
	fmt.Scan(&l1, &r1, &l2, &r2)
	fmt.Println(max(min(r1, r2)-max(l1, l2), 0))
}

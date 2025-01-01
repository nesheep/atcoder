package main

import (
	"fmt"
)

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	m := make(map[int]int, 4)
	m[a]++
	m[b]++
	m[c]++
	m[d]++

	if len(m) == 2 {
		fmt.Println("Yes")
		return
	}

	fmt.Println("No")
}

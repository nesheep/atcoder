package main

import (
	"fmt"
	"os"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	m := make(map[int]struct{}, 3)
	m[a] = struct{}{}
	m[b] = struct{}{}

	if len(m) != 2 {
		fmt.Println(-1)
		os.Exit(0)
	}

	for i := 1; i <= 3; i++ {
		if _, ok := m[i]; !ok {
			fmt.Println(i)
			os.Exit(0)
		}
	}
}

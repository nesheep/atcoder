package main

import "fmt"

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)

	c := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&c[i])
	}

	for i := range c {
		if c[i] == a+b {
			fmt.Println(i + 1)
			return
		}
	}
}

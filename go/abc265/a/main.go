package main

import (
	"fmt"
	"os"
)

func main() {
	var x, y, n int
	fmt.Scan(&x, &y, &n)

	if x*3 <= y {
		fmt.Println(x * n)
		os.Exit(0)
	}

	fmt.Println(y*(n/3) + x*(n%3))
}

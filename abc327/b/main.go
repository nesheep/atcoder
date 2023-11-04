package main

import (
	"fmt"
	"os"
)

func main() {
	var b int
	fmt.Scan(&b)

	for i := 1; i < 16; i++ {
		a := i
		for j := 1; j < i; j++ {
			a *= i
		}
		if a == b {
			fmt.Println(i)
			os.Exit(0)
		}
	}

	fmt.Println(-1)
}

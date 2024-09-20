package main

import (
	"fmt"
	"os"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	for i := 0; i <= 9; i++ {
		if i != a+b {
			fmt.Println(i)
			os.Exit(0)
		}
	}
}

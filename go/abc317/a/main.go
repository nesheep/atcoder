package main

import (
	"fmt"
	"os"
)

func main() {
	var n, h, x int
	fmt.Scan(&n, &h, &x)

	for i := 1; i <= n; i++ {
		var p int
		fmt.Scan(&p)
		if h+p >= x {
			fmt.Println(i)
			os.Exit(0)
		}
	}
}

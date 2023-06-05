package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	ans := 1
	for i := 0; i < b; i++ {
		ans *= a
	}
	fmt.Println(ans)
}

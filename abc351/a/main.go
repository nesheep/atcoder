package main

import "fmt"

func main() {
	var a_total, b_total int

	for i := 0; i < 9; i++ {
		var a int
		fmt.Scan(&a)
		a_total += a
	}

	for i := 0; i < 8; i++ {
		var b int
		fmt.Scan(&b)
		b_total += b
	}

	fmt.Println(a_total - b_total + 1)
}

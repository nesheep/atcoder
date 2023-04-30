package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	for a >= 1 && b >= 1 {
		if a < b {
			b = b % a
		} else {
			a = a % b
		}
	}

	if a >= 1 {
		fmt.Println(a)
	} else {
		fmt.Println(b)
	}
}

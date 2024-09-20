package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	f := 1
	for i := 2; i <= n; i++ {
		f *= i
	}

	fmt.Println(f)
}

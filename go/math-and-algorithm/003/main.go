package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var a int
	var sum int
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		sum += a
	}

	fmt.Println(sum)
}

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	switch {
	case n < 1000:
		fmt.Println(n)
	case n < 10000:
		fmt.Println(n - n%10)
	case n < 100000:
		fmt.Println(n - n%100)
	case n < 1000000:
		fmt.Println(n - n%1000)
	case n < 10000000:
		fmt.Println(n - n%10000)
	case n < 100000000:
		fmt.Println(n - n%100000)
	default:
		fmt.Println(n - n%1000000)
	}
}

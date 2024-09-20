package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	m := 998244353
	fmt.Println((n%m + m) % m)
}

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Printf("%02d\n", n%100)
}

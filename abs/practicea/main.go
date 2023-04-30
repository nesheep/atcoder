package main

import "fmt"

var (
	a, b, c int
	s       string
)

func main() {
	fmt.Scan(&a, &b, &c, &s)
	fmt.Printf("%d %s\n", a+b+c, s)
}

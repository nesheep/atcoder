package main

import "fmt"

func main() {
	var a, b, c, d, e int
	fmt.Scan(&a, &b, &c, &d, &e)
	m := make(map[int]struct{}, 5)
	m[a] = struct{}{}
	m[b] = struct{}{}
	m[c] = struct{}{}
	m[d] = struct{}{}
	m[e] = struct{}{}
	fmt.Println(len(m))
}

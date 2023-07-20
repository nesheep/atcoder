package main

import "fmt"

func main() {
	var l, r int
	fmt.Scan(&l, &r)

	atcoder := []byte{'a', 't', 'c', 'o', 'd', 'e', 'r'}
	fmt.Printf("%s\n", atcoder[l-1:r])
}

package main

import "fmt"

func main() {
	var k int
	fmt.Scan(&k)

	s := make([]byte, 0, k)
	for i := 0; i < k; i++ {
		s = append(s, 'A'+byte(i))
	}

	fmt.Printf("%s\n", s)
}

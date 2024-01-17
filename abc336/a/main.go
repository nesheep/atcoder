package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	s := make([]byte, n+3)
	s[0] = 'L'
	s[n+1] = 'n'
	s[n+2] = 'g'
	for i := 1; i <= n; i++ {
		s[i] = 'o'
	}
	fmt.Printf("%s\n", s)
}

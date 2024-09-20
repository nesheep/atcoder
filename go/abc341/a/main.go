package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	s := make([]byte, n*2+1)
	for i := range s {
		if i%2 == 0 {
			s[i] = '1'
		} else {
			s[i] = '0'
		}
	}

	fmt.Println(string(s))
}

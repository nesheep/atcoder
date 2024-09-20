package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	s := make([]byte, n)
	for i := 0; i < n; i++ {
		if (i+1)%3 == 0 {
			s[i] = 'x'
		} else {
			s[i] = 'o'
		}
	}

	fmt.Println(string(s))
}

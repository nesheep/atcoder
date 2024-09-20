package main

import "fmt"

func main() {
	s := make([][]byte, 10)
	for i := range s {
		fmt.Scan(&s[i])
	}

	a, c := 10, 10
	b, d := -1, -1
	for i := range s {
		for j := range s[i] {
			if s[i][j] == '#' {
				if i < a {
					a = i
				}
				if i > b {
					b = i
				}
				if j < c {
					c = j
				}
				if j > d {
					d = j
				}
			}
		}
	}

	fmt.Println(a+1, b+1)
	fmt.Println(c+1, d+1)
}

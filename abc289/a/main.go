package main

import "fmt"

func main() {
	var s []byte
	fmt.Scan(&s)

	for i := range s {
		if s[i] == '0' {
			s[i] = '1'
		} else {
			s[i] = '0'
		}
	}

	fmt.Printf("%s\n", s)
}

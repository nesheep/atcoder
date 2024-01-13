package main

import "fmt"

func main() {
	var s []byte
	fmt.Scan(&s)
	s[len(s)-1] = '4'
	fmt.Printf("%s\n", s)
}

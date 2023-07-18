package main

import "fmt"

func main() {
	var s []byte
	fmt.Scan(&s)
	fmt.Printf("%c\n", s[len(s)/2])
}

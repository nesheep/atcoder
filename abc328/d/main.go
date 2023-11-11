package main

import (
	"bytes"
	"fmt"
)

func main() {
	var s []byte
	fmt.Scan(&s)

	for {
		l := len(s)
		s = bytes.Replace(s, []byte("ABC"), []byte(""), -1)
		if l == len(s) {
			break
		}
	}

	fmt.Printf("%s\n", s)
}

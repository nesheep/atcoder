package main

import (
	"bytes"
	"fmt"
)

func main() {
	var s []byte
	fmt.Scan(&s)

	var cnt int
	for _, v := range s {
		if v < 'a' {
			cnt++
		}
	}

	if cnt*2 > len(s) {
		fmt.Printf("%s\n", bytes.ToUpper(s))
	} else {
		fmt.Printf("%s\n", bytes.ToLower(s))
	}
}

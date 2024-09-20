package main

import (
	"bytes"
	"fmt"
)

func main() {
	var n int
	var s []byte

	fmt.Scan(&n, &s)
	if bytes.Contains(s, []byte("ab")) || bytes.Contains(s, []byte("ba")) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

package main

import (
	"bytes"
	"fmt"
)

func main() {
	var n, m int
	var s, t []byte
	fmt.Scan(&n, &m, &s, &t)

	isPrefix := bytes.HasPrefix(t, s)
	isSuffix := bytes.HasSuffix(t, s)

	switch {
	case isPrefix && isSuffix:
		fmt.Println(0)
	case isPrefix:
		fmt.Println(1)
	case isSuffix:
		fmt.Println(2)
	default:
		fmt.Println(3)
	}
}

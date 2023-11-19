package main

import (
	"bytes"
	"fmt"
)

func main() {
	var s []byte
	fmt.Scan(&s)

	t := make([][]byte, len(s))
	for i := range s {
		t[i] = []byte{s[i]}
	}

	fmt.Printf("%s\n", bytes.Join(t, []byte{' '}))
}

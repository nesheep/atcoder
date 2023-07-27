package main

import (
	"fmt"
	"os"
)

func main() {
	var s []byte
	fmt.Scan(&s)

	m := make(map[byte]int, len(s))
	for _, c := range s {
		m[c]++
	}

	for k, v := range m {
		if v == 1 {
			fmt.Printf("%c\n", k)
			os.Exit(0)
		}
	}
	fmt.Println(-1)
}

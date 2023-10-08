package main

import (
	"fmt"
	"os"
)

func main() {
	var n int
	var s []byte
	fmt.Scan(&n, &s)

	for i := 0; i < n-2; i++ {
		if s[i] == 'A' && s[i+1] == 'B' && s[i+2] == 'C' {
			fmt.Println(i + 1)
			os.Exit(0)
		}
	}

	fmt.Println(-1)
}

package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "wbwbwwbwbwbw"

	var w, b int
	fmt.Scan(&w, &b)

	ss := strings.Repeat(s, 18)
	for i := 0; i < len(s); i++ {
		var wc, bc int
		for j := 0; j < w+b; j++ {
			if ss[i+j] == 'w' {
				wc++
			} else {
				bc++
			}
		}
		if wc == w && bc == b {
			fmt.Println("Yes")
			return
		}
	}

	fmt.Println("No")
}

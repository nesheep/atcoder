package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rd := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(rd, &n)

	a := make([][]byte, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &a[i])
	}

	b := make([][]byte, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &b[i])
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if a[i][j] != b[i][j] {
				fmt.Println(i+1, j+1)
				os.Exit(0)
			}
		}
	}
}

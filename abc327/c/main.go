package main

import (
	"fmt"
	"os"
)

func jadge(b [9]int) bool {
	m := make(map[int]bool, 9)
	for _, v := range b {
		m[v] = true
	}
	return len(m) == 9
}

func main() {
	var a [9][9]int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Scan(&a[i][j])
		}
	}

	for i := 0; i < 9; i++ {
		if !jadge(a[i]) {
			fmt.Println("No")
			os.Exit(0)
		}
	}

	for j := 0; j < 9; j++ {
		var b [9]int
		for i := 0; i < 9; i++ {
			b[i] = a[i][j]
		}
		if !jadge(b) {
			fmt.Println("No")
			os.Exit(0)
		}
	}

	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			var b [9]int
			for k := 0; k < 3; k++ {
				for l := 0; l < 3; l++ {
					b[k*3+l] = a[i+k][j+l]
				}
			}
			if !jadge(b) {
				fmt.Println("No")
				os.Exit(0)
			}
		}
	}

	fmt.Println("Yes")
}

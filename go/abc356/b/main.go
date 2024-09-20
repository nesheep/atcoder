package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rd := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(rd, &n, &m)

	a := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(rd, &a[i])
	}

	x := make([][]int, n)
	for i := 0; i < n; i++ {
		x[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Fscan(rd, &x[i][j])
		}
	}

	b := make([]int, m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			b[j] += x[i][j]
		}
	}

	for i := 0; i < m; i++ {
		if b[i] < a[i] {
			fmt.Println("No")
			os.Exit(0)
		}
	}

	fmt.Println("Yes")
}

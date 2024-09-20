package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var n int
	fmt.Fscan(rd, &n)

	a := make([][]byte, n)
	for i := range a {
		fmt.Fscan(rd, &a[i])
	}

	b := make([][]byte, n)
	for i := range b {
		b[i] = make([]byte, n)
	}

	b[0][0] = a[1][0]
	for i := 1; i < n; i++ {
		b[0][i] = a[0][i-1]
	}

	for i := 1; i < n-1; i++ {
		b[i][0] = a[i+1][0]
		for j := 1; j < n-1; j++ {
			b[i][j] = a[i][j]
		}
		b[i][n-1] = a[i-1][n-1]
	}

	for i := 0; i < n-1; i++ {
		b[n-1][i] = a[n-1][i+1]
	}
	b[n-1][n-1] = a[n-2][n-1]

	for i := range b {
		fmt.Fprintf(wr, "%s\n", b[i])
	}
}

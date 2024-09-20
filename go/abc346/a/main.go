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

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &a[i])
	}

	b := make([]any, n-1)
	for i := 0; i < n-1; i++ {
		b[i] = a[i] * a[i+1]
	}

	fmt.Fprintln(wr, b...)
}

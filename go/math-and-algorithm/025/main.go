package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n int
	fmt.Fscan(r, &n)

	a := make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &a[i])
	}

	b := make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &b[i])
	}

	var e float64
	for i := 0; i < n; i++ {
		e += a[i]/3 + b[i]*2/3
	}

	fmt.Fprintf(w, "%.12f\n", e)
}

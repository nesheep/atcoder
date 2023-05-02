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

	var p, q, e float64
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &p, &q)
		e += q / p
	}

	fmt.Fprintf(w, "%.12f\n", e)
}

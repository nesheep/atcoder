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

	var b, bsum float64
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &b)
		bsum += b
	}

	var c, csum float64
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &c)
		csum += c
	}

	ans := bsum/float64(n) + csum/float64(n)
	fmt.Fprintf(w, "%.12f\n", ans)
}

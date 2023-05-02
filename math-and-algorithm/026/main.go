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

	var e float64
	for i := 1; i <= n; i++ {
		e += float64(n) / float64(i)
	}

	fmt.Fprintf(w, "%.12f\n", e)
}

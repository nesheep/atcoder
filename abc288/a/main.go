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

	var a, b int
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &a, &b)
		fmt.Fprintln(w, a+b)
	}
}

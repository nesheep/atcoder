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

	var d int
	m := make(map[int]bool, 100)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &d)
		m[d] = true
	}

	fmt.Fprintln(w, len(m))
}

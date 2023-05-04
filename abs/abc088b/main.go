package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n int
	fmt.Fscan(r, &n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &a[i])
	}

	sort.Ints(a)

	var ac, bb int
	for i := 1; i <= n; i++ {
		if i%2 == 1 {
			ac += a[n-i]
		} else {
			bb += a[n-i]
		}
	}

	fmt.Fprintln(w, ac-bb)
}

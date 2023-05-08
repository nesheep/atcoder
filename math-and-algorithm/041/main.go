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

	var t, n int
	fmt.Fscan(r, &t, &n)

	var tl, tr int
	as := make([]int, t+1)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &tl, &tr)
		as[tl]++
		as[tr]--
	}

	var ans int
	for i := 0; i < t; i++ {
		ans += as[i]
		fmt.Fprintln(w, ans)
	}
}

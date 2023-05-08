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

	ls := make([]int, n)
	rs := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &ls[i], &rs[i])
	}

	as := make([]int, t+1)
	for i := 0; i < n; i++ {
		as[ls[i]]++
		as[rs[i]]--
	}

	ts := make([]int, t)
	ts[0] = as[0]
	for i := 1; i < t; i++ {
		ts[i] = ts[i-1] + as[i]
	}

	for i := 0; i < t; i++ {
		fmt.Fprintln(w, ts[i])
	}
}

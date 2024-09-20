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

	var n, q int
	fmt.Fscan(r, &n, &q)

	as := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(r, &as[i])
	}

	ls := make([]int, q)
	rs := make([]int, q)
	for i := 0; i < q; i++ {
		fmt.Fscan(r, &ls[i], &rs[i])
	}

	bs := make([]int, n+1)
	for i := 1; i <= n; i++ {
		bs[i] = bs[i-1] + as[i]
	}

	for i := 0; i < q; i++ {
		fmt.Fprintln(w, bs[rs[i]]-bs[ls[i]-1])
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	var n, q int
	fmt.Fscan(r, &n, &q)

	ls := make([]int, q+1)
	rs := make([]int, q+1)
	xs := make([]int, q+1)
	for i := 1; i <= q; i++ {
		fmt.Fscan(r, &ls[i], &rs[i], &xs[i])
	}

	as := make([]int, n+2)
	for i := 1; i <= q; i++ {
		as[ls[i]] += xs[i]
		as[rs[i]+1] -= xs[i]
	}

	ans := make([]string, 0, n-1)
	for i := 2; i <= n; i++ {
		if as[i] < 0 {
			ans = append(ans, ">")
		} else if as[i] == 0 {
			ans = append(ans, "=")
		} else {
			ans = append(ans, "<")
		}
	}

	fmt.Println(strings.Join(ans, ""))
}

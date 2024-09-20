package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rd := bufio.NewReader(os.Stdin)

	var n, q int
	var s []byte
	fmt.Fscan(rd, &n, &q, &s)

	t := make([]int, n)
	for i := 0; i < n-1; i++ {
		t[i+1] = t[i]
		if s[i] == s[i+1] {
			t[i+1]++
		}
	}

	for i := 0; i < q; i++ {
		var l, r int
		fmt.Fscan(rd, &l, &r)
		fmt.Println(t[r-1] - t[l-1])
	}
}

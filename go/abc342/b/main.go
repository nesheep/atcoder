package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rd := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(rd, &n)

	p := make(map[int]int, n)
	for i := 1; i <= n; i++ {
		var pi int
		fmt.Fscan(rd, &pi)
		p[pi] = i
	}

	var q int
	fmt.Fscan(rd, &q)

	for i := 0; i < q; i++ {
		var a, b int
		fmt.Fscan(rd, &a, &b)
		if p[a] < p[b] {
			fmt.Println(a)
		} else {
			fmt.Println(b)
		}
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var n, x int
	fmt.Fscan(rd, &n, &x)

	a := make(map[int]struct{})
	var v int
	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &v)
		a[v] = struct{}{}
	}

	for k := range a {
		if _, ok := a[x+k]; ok {
			fmt.Fprintln(wr, "Yes")
			return
		}
	}
	fmt.Fprintln(wr, "No")
}

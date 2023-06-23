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

	var n int
	fmt.Fscan(rd, &n)

	var ans int
	var v int
	a := make(map[int]int)
	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &v)
		a[v]++
		if a[v]%2 == 0 {
			ans++
		}
	}

	fmt.Fprintln(wr, ans)
}

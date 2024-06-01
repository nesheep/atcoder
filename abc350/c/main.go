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

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &a[i])
	}

	ans := make([][2]int, 0, n)

	for i := 0; i < n; i++ {
		for a[i] != i+1 {
			j := a[i] - 1
			a[i], a[j] = a[j], a[i]
			ans = append(ans, [2]int{i, j})
		}
	}

	fmt.Fprintln(wr, len(ans))
	for _, v := range ans {
		fmt.Fprintln(wr, v[0]+1, v[1]+1)
	}
}

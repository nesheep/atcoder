package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var n int
	fmt.Fscan(rd, &n)

	a := make([]map[int]bool, n)
	for i := 0; i < n; i++ {
		var c int
		fmt.Fscan(rd, &c)
		a[i] = make(map[int]bool, c)
		for j := 0; j < c; j++ {
			var v int
			fmt.Fscan(rd, &v)
			a[i][v] = true
		}
	}

	var x int
	fmt.Fscan(rd, &x)

	m := 38
	b := []int{}
	for i := range a {
		if !a[i][x] {
			continue
		}
		if len(a[i]) > m {
			continue
		}
		if len(a[i]) == m {
			b = append(b, i)
			continue
		}
		m = len(a[i])
		b = []int{i}
	}

	ans := make([]string, len(b))
	for i := range b {
		ans[i] = strconv.Itoa(b[i] + 1)
	}

	fmt.Fprintln(wr, len(ans))
	fmt.Fprintln(wr, strings.Join(ans, " "))
}

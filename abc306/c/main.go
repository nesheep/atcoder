package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var n int
	fmt.Fscan(rd, &n)

	a := make([]int, n*3)
	for i := 0; i < n*3; i++ {
		fmt.Fscan(rd, &a[i])
	}

	e := make([]bool, n)
	f := make([]int, n)
	for i := range f {
		f[i] = -1
	}

	for i := range a {
		if !e[a[i]-1] {
			e[a[i]-1] = true
		} else if f[a[i]-1] == -1 {
			f[a[i]-1] = i
		}
	}

	sort.Ints(f)

	ans := make([]string, 0, n)
	for i := range f {
		ans = append(ans, fmt.Sprint(a[f[i]]))
	}

	fmt.Fprintln(wr, strings.Join(ans, " "))
}

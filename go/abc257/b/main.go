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

	var n, k, q int
	fmt.Fscan(rd, &n, &k, &q)

	a := make(map[int]int, k)
	for i := 1; i <= k; i++ {
		var ai int
		fmt.Fscan(rd, &ai)
		a[i] = ai
	}

	for i := 0; i < q; i++ {
		var l int
		fmt.Fscan(rd, &l)
		if a[l] == n || a[l+1] == a[l]+1 {
			continue
		}
		a[l]++
	}

	s := make([]string, 0, k)
	for i := 1; i <= k; i++ {
		s = append(s, strconv.Itoa(a[i]))
	}

	fmt.Fprintln(wr, strings.Join(s, " "))
}

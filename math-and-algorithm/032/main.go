package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n, x int
	fmt.Fscan(r, &n, &x)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &a[i])
	}

	sort.Ints(a)

	if BinarySearch(a, x) {
		fmt.Fprintln(w, "Yes")
	} else {
		fmt.Fprintln(w, "No")
	}
}

func BinarySearch(a []int, x int) bool {
	l := 0
	r := len(a)
	for l < r {
		m := (l + r) / 2
		if a[m] == x {
			return true
		}
		if a[m] > x {
			r = m
		} else {
			l = m + 1
		}
	}
	return false
}

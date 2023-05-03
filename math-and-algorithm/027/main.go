package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n int
	fmt.Fscan(r, &n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &a[i])
	}

	MargeSort(a)

	ans := make([]string, 0, n)
	for _, v := range a {
		ans = append(ans, strconv.Itoa(v))
	}

	fmt.Fprintln(w, strings.Join(ans, " "))
}

func MargeSort(a []int) {
	c := make([]int, len(a))

	var f func(l, r int)
	f = func(l, r int) {
		if r-l == 1 {
			return
		}

		m := (l + r) / 2
		f(l, m)
		f(m, r)

		c1 := l
		c2 := m
		var cnt int
		for c1 != m || c2 != r {
			if c1 == m {
				c[cnt] = a[c2]
				c2++
			} else if c2 == r {
				c[cnt] = a[c1]
				c1++
			} else {
				if a[c1] < a[c2] {
					c[cnt] = a[c1]
					c1++
				} else {
					c[cnt] = a[c2]
					c2++
				}
			}
			cnt++
		}

		for i := 0; i < cnt; i++ {
			a[l+i] = c[i]
		}
	}

	f(0, len(a))
}

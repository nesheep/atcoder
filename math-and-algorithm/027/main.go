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
	if len(a) == 1 {
		return
	}

	m := len(a) / 2
	MargeSort(a[:m])
	MargeSort(a[m:])

	c1 := 0
	c2 := m
	cnt := 0
	c := make([]int, len(a))
	for c1 != m || c2 != len(a) {
		if c1 == m {
			c[cnt] = a[c2]
			c2++
		} else if c2 == len(a) {
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

	copy(a, c)
}

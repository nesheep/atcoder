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

	var n, m int
	fmt.Fscan(rd, &n, &m)

	a := make(map[int]bool, m)
	var tmp int
	for i := 0; i < m; i++ {
		fmt.Fscan(rd, &tmp)
		a[tmp] = true
	}

	ans := make([]string, 0, n)
	c := []int{}
	for i := 1; i <= n; i++ {
		if a[i] {
			c = append(c, i)
		} else {
			ans = append(ans, strconv.Itoa(i))
			for j := len(c) - 1; j >= 0; j-- {
				ans = append(ans, strconv.Itoa(c[j]))
			}
			c = []int{}
		}
	}

	fmt.Fprintln(wr, strings.Join(ans, " "))
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	rd := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(rd, &n, &m)

	a := make([]int, n)
	for i := range a {
		fmt.Fscan(rd, &a[i])
	}

	b := make([]int, m)
	for i := range b {
		fmt.Fscan(rd, &b[i])
	}

	c := make([]int, 0, n+m)
	c = append(c, a...)
	c = append(c, b...)
	sort.Ints(c)

	am := make(map[int]struct{}, n)
	for i := range a {
		am[a[i]] = struct{}{}
	}

	for i := 1; i < len(c); i++ {
		_, ok1 := am[c[i-1]]
		_, ok2 := am[c[i]]
		if ok1 && ok2 {
			fmt.Println("Yes")
			os.Exit(0)
		}
	}

	fmt.Println("No")
}

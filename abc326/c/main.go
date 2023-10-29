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
	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &a[i])
	}

	sort.Ints(a)

	b := make([]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			if a[i]-a[j] < m {
				b[j]++
			}
		}
	}

	var ans int
	for i := 0; i < n; i++ {
		if b[i] > ans {
			ans = b[i]
		}
	}

	fmt.Println(ans)
}

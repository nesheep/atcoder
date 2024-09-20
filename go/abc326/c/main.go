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

	var ans int
	for i := 0; i < n; i++ {
		j := sort.Search(n, func(k int) bool { return a[k] >= a[i]+m })
		d := j - i
		if d > ans {
			ans = d
		}
	}

	fmt.Println(ans)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	rd := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscan(rd, &n, &k)

	var sum int
	m := make([][2]int, n)
	for i := range m {
		fmt.Fscan(rd, &m[i][0], &m[i][1])
		sum += m[i][1]
	}

	sort.Slice(m, func(i, j int) bool {
		return m[i][0] < m[j][0]
	})

	if sum <= k {
		fmt.Println(1)
		return
	}

	l := sum
	for _, v := range m {
		l -= v[1]
		if l <= k {
			fmt.Println(v[0] + 1)
			return
		}
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	rd := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(rd, &n)

	ls := make(map[int]int, n)
	rs := make(map[int]int, n)

	m := make(map[int]struct{}, n*2)

	for i := 0; i < n; i++ {
		var l, r int
		fmt.Fscan(rd, &l, &r)

		ls[l]++
		rs[r]++

		m[l] = struct{}{}
		m[r] = struct{}{}
	}

	s := make([]int, 0, len(m))
	for k := range m {
		s = append(s, k)
	}
	sort.Ints(s)

	var cnt, ans int
	for _, i := range s {
		lc := ls[i]
		ans += (cnt + (cnt + lc - 1)) * lc / 2
		cnt += lc - rs[i]
	}

	fmt.Println(ans)
}

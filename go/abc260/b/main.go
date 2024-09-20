package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type person struct {
	id, a, b int
}

func main() {
	rd := bufio.NewReader(os.Stdin)

	var n, x, y, z int
	fmt.Fscan(rd, &n, &x, &y, &z)

	p := make([]person, n)
	for i := range p {
		p[i].id = i + 1
		fmt.Fscan(rd, &p[i].a)
	}
	for i := range p {
		fmt.Fscan(rd, &p[i].b)
	}

	ans := make([]bool, n+1)
	tmp := make([]person, n)

	copy(tmp, p)
	sort.SliceStable(tmp, func(i, j int) bool {
		return tmp[i].a > tmp[j].a
	})

	for i, cnt := 0, 0; cnt < x; i++ {
		if !ans[tmp[i].id] {
			ans[tmp[i].id] = true
			cnt++
		}
	}

	copy(tmp, p)
	sort.SliceStable(tmp, func(i, j int) bool {
		return tmp[i].b > tmp[j].b
	})

	for i, cnt := 0, 0; cnt < y; i++ {
		if !ans[tmp[i].id] {
			ans[tmp[i].id] = true
			cnt++
		}
	}

	copy(tmp, p)
	sort.SliceStable(tmp, func(i, j int) bool {
		return tmp[i].a+tmp[i].b > tmp[j].a+tmp[j].b
	})

	for i, cnt := 0, 0; cnt < z; i++ {
		if !ans[tmp[i].id] {
			ans[tmp[i].id] = true
			cnt++
		}
	}

	for i := 1; i <= n; i++ {
		if ans[i] {
			fmt.Println(i)
		}
	}
}

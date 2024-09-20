package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rd := bufio.NewReader(os.Stdin)

	var n, w int
	fmt.Fscan(rd, &n, &w)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &a[i])
	}

	m := make(map[int]struct{})
	for i := 0; i < n; i++ {
		s := a[i]
		if s <= w {
			m[s] = struct{}{}
		}
		for j := i + 1; j < n; j++ {
			s := a[i] + a[j]
			if s <= w {
				m[s] = struct{}{}
			}
			for k := j + 1; k < n; k++ {
				s := a[i] + a[j] + a[k]
				if s <= w {
					m[s] = struct{}{}
				}
			}
		}
	}

	fmt.Println(len(m))
}

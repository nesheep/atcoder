package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rd := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(rd, &n, &m)

	s := make([]map[int]bool, m)

	for i := range s {
		var c int
		fmt.Fscan(rd, &c)
		s[i] = make(map[int]bool, c)
		for j := 0; j < c; j++ {
			var d int
			fmt.Fscan(rd, &d)
			s[i][d] = true
		}
	}

	var ans int
	for i := 0; i < 1<<m; i++ {
		t := make(map[int]bool)
		for j := 0; j < m; j++ {
			if i>>j&1 != 1 {
				continue
			}
			for k := range s[j] {
				if s[j][k] {
					t[k] = true
				}
			}
		}
		if len(t) == n {
			ans++
		}
	}

	fmt.Println(ans)
}

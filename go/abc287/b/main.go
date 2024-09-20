package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(r, &n, &m)

	ss := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &ss[i])
	}

	ts := make([]string, m)
	for j := 0; j < m; j++ {
		fmt.Fscan(r, &ts[j])
	}

	var cnt int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if ss[i][3:] == ts[j] {
				cnt++
				break
			}
		}
	}

	fmt.Println(cnt)
}

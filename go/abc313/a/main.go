package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rd := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(rd, &n)

	p := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &p[i])
	}

	m := 0
	for i := 1; i < n; i++ {
		if p[i] > m {
			m = p[i]
		}
	}

	ans := 0
	if p[0] <= m {
		ans = m - p[0] + 1
	}

	fmt.Println(ans)
}

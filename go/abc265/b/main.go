package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rd := bufio.NewReader(os.Stdin)

	var n, m, t int
	fmt.Fscan(rd, &n, &m, &t)

	a := make([]int, n-1)
	for i := range a {
		fmt.Fscan(rd, &a[i])
	}

	b := make(map[int]int, m)
	for i := 0; i < m; i++ {
		var x, y int
		fmt.Fscan(rd, &x, &y)
		b[x-1] = y
	}

	ans := "Yes"
	for i := range a {
		t += b[i]
		t -= a[i]
		if t <= 0 {
			ans = "No"
			break
		}
	}

	fmt.Println(ans)
}

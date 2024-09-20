package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(r, &n)

	as := make([]int, n)
	for i := 1; i <= n-1; i++ {
		fmt.Fscan(r, &as[i])
	}

	var m int
	fmt.Fscan(r, &m)

	bs := make([]int, m+1)
	for i := 1; i <= m; i++ {
		fmt.Fscan(r, &bs[i])
	}

	cs := make([]int, n+1)
	for i := 2; i <= n; i++ {
		cs[i] = cs[i-1] + as[i-1]
	}

	var ans int
	for i := 2; i <= m; i++ {
		ans += abs(cs[bs[i]] - cs[bs[i-1]])
	}

	fmt.Println(ans)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

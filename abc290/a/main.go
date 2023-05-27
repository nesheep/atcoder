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

	as := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(r, &as[i])
	}

	var b, sum int
	for i := 0; i < m; i++ {
		fmt.Fscan(r, &b)
		sum += as[b]
	}

	fmt.Println(sum)
}

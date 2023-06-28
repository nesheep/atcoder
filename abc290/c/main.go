package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rd := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscan(rd, &n, &k)

	a := make(map[int]bool, n)
	var v int
	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &v)
		a[v] = true
	}

	var m int
	for i := 0; i < k; i++ {
		if !a[i] {
			break
		}
		m++
	}

	fmt.Println(m)
}

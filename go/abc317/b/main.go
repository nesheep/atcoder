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

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &a[i])
	}

	sort.Ints(a)

	for i := 1; i < n; i++ {
		if a[i] == a[i-1]+1 {
			continue
		}
		fmt.Println(a[i] - 1)
		os.Exit(0)
	}
}

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

	a := make([]int, n)
	for i := range a {
		fmt.Fscan(rd, &a[i])
	}

	var m int
	for i := range a {
		if a[i] > m {
			m = a[i]
		}
	}

	var ans int
	for i := range a {
		if a[i] != m && a[i] > ans {
			ans = a[i]
		}
	}

	fmt.Println(ans)
}

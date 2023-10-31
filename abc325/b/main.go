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

	t := make([]int, 24)
	for i := 0; i < n; i++ {
		var w, x int
		fmt.Fscan(rd, &w, &x)
		s := 24 - x
		for j := 9; j < 18; j++ {
			t[(s+j)%24] += w
		}
	}

	var m int
	for i := range t {
		if t[i] > m {
			m = t[i]
		}
	}

	fmt.Println(m)
}

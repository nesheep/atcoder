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

	a := make([]int, n)
	for i := range a {
		fmt.Fscan(rd, &a[i])
	}

	b := make(map[int]bool, k)
	for i := 0; i < k; i++ {
		var t int
		fmt.Fscan(rd, &t)
		b[t] = true
	}

	var dMax int
	for _, v := range a {
		if v > dMax {
			dMax = v
		}
	}

	var dMaxes []int
	for i, v := range a {
		if v == dMax {
			dMaxes = append(dMaxes, i+1)
		}
	}

	for _, v := range dMaxes {
		if b[v] {
			fmt.Println("Yes")
			os.Exit(0)
		}
	}

	fmt.Println("No")
}

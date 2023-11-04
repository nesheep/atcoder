package main

import (
	"bufio"
	"fmt"
	"os"
)

func contains(s []int, v int) bool {
	for _, vv := range s {
		if vv == v {
			return true
		}
	}
	return false
}

func main() {
	rd := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(rd, &n, &m)

	a := make([]int, m)
	for i := range a {
		fmt.Fscan(rd, &a[i])
	}

	b := make([]int, m)
	for i := range b {
		fmt.Fscan(rd, &b[i])
	}

	x := make([][]int, n+1)
	for i := range a {
		x[a[i]] = append(x[a[i]], b[i])
		x[b[i]] = append(x[b[i]], a[i])
	}

	for i := range x {
		for j := range x[i] {
			for k := range x[i] {
				if contains(x[x[i][j]], x[i][k]) {
					fmt.Println("No")
					os.Exit(0)
				}
			}
		}
	}

	fmt.Println("Yes")
}

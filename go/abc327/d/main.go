package main

import (
	"bufio"
	"fmt"
	"os"
)

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

	x := make([]map[int]bool, n+1)
	for i := range a {
		if x[a[i]] == nil {
			x[a[i]] = make(map[int]bool)
		}
		if x[b[i]] == nil {
			x[b[i]] = make(map[int]bool)
		}
		x[a[i]][b[i]] = true
		x[b[i]][a[i]] = true
	}

	for i := range x {
		for j := range x[i] {
			for k := range x[i] {
				if x[j][k] {
					fmt.Println("No")
					os.Exit(0)
				}
			}
		}
	}

	fmt.Println("Yes")
}

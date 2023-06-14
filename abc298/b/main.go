package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var n int
	fmt.Fscan(rd, &n)

	a := make([][]bool, n)
	for i := range a {
		a[i] = make([]bool, n)
		for j := range a[i] {
			fmt.Fscan(rd, &a[i][j])
		}
	}

	b := make([][]bool, n)
	for i := range b {
		b[i] = make([]bool, n)
		for j := range b[i] {
			fmt.Fscan(rd, &b[i][j])
		}
	}

	ok0, ok1, ok2, ok3 := true, true, true, true
	for i := range a {
		for j := range a[i] {
			if a[i][j] && !b[i][j] {
				ok0 = false
			}
			if a[n-j-1][i] && !b[i][j] {
				ok1 = false
			}
			if a[n-i-1][n-j-1] && !b[i][j] {
				ok2 = false
			}
			if a[j][n-i-1] && !b[i][j] {
				ok3 = false
			}
		}
	}

	if ok0 || ok1 || ok2 || ok3 {
		fmt.Fprintln(wr, "Yes")
	} else {
		fmt.Fprintln(wr, "No")
	}
}

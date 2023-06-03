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

	var n, m int
	fmt.Fscan(rd, &n, &m)

	as := make([][]int, m)
	for i := 0; i < m; i++ {
		a := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(rd, &a[j])
		}
		as[i] = a
	}

	ps := make([][]bool, n)
	for i := 0; i < n; i++ {
		p := make([]bool, n)
		ps[i] = p
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n-1; j++ {
			if as[i][j] < as[i][j+1] {
				ps[as[i][j]-1][as[i][j+1]-1] = true
			} else {
				ps[as[i][j+1]-1][as[i][j]-1] = true
			}
		}
	}

	var cnt int
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if !ps[i][j] {
				cnt++
			}
		}
	}

	fmt.Fprintln(wr, cnt)
}

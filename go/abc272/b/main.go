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

	a := make([][]bool, n)
	for i := range a {
		a[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		var k int
		fmt.Fscan(rd, &k)
		x := make([]int, k)
		for j := 0; j < k; j++ {
			fmt.Fscan(rd, &x[j])
			x[j]--
		}
		for j := 0; j < k; j++ {
			for l := 0; l < k; l++ {
				a[x[j]][x[l]] = true
				a[x[l]][x[j]] = true
			}
		}
	}

	ans := "Yes"
Loop:
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i != j && !a[i][j] {
				ans = "No"
				break Loop
			}
		}
	}

	fmt.Println(ans)
}

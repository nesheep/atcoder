package main

import "fmt"

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)

	ans := make([][]byte, a*n)
	for i := range ans {
		ans[i] = make([]byte, b*n)
	}

	for i := range ans {
		for j := range ans[i] {
			if (i/a+j/b)%2 == 1 {
				ans[i][j] = '#'
			} else {
				ans[i][j] = '.'
			}
		}
	}

	for i := range ans {
		fmt.Printf("%s\n", ans[i])
	}
}

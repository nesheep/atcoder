package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	a := make([][]int, n)
	a[0] = []int{1}
	for i := 1; i < n; i++ {
		a[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				a[i][j] = 1
			} else {
				a[i][j] = a[i-1][j-1] + a[i-1][j]
			}
		}
	}

	for i := 0; i < n; i++ {
		b := make([]string, len(a[i]))
		for j := 0; j < len(a[i]); j++ {
			b[j] = strconv.Itoa(a[i][j])
		}
		fmt.Println(strings.Join(b, " "))
	}
}

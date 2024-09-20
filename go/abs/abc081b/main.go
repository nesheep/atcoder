package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	min := math.MaxInt64
	for i := range a {
		var cnt int
		tmp := a[i]
		for tmp%2 == 0 {
			tmp /= 2
			cnt++
		}
		if cnt < min {
			min = cnt
		}
	}

	fmt.Println(min)
}

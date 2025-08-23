package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	rs := make([]int, n)
	cs := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&rs[i], &cs[i])
	}

	d := max([]int{
		max(rs) - min(rs),
		max(cs) - min(cs),
	})

	ans := (d + 1) / 2

	fmt.Println(ans)
}

func min(s []int) int {
	x := s[0]
	for _, v := range s {
		if v < x {
			x = v
		}
	}
	return x
}

func max(s []int) int {
	x := s[0]
	for _, v := range s {
		if v > x {
			x = v
		}
	}
	return x
}

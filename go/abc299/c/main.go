package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var n int
	var s []byte
	fmt.Scan(&n, &s)

	var ans, x int
	for i := range s {
		if s[i] == 'o' {
			x++
		} else {
			x = 0
		}
		ans = max(ans, x)
	}

	if ans == 0 || ans == n {
		ans = -1
	}
	fmt.Println(ans)
}

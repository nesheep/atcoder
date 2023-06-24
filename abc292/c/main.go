package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	dc := make([]int, n+1)
	for i := 1; i <= n; i++ {
		for j := i; j <= n; j += i {
			dc[j]++
		}
	}

	var ans int
	for i := 1; i < n; i++ {
		ans += dc[i] * dc[n-i]
	}

	fmt.Println(ans)
}

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		for j := i; j <= n; j += i {
			f[j]++
		}
	}

	var ans int
	for i := 1; i <= n; i++ {
		ans += i * f[i]
	}

	fmt.Println(ans)
}

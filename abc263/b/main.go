package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	p := make(map[int]int, n-1)
	for i := 2; i <= n; i++ {
		var pi int
		fmt.Scan(&pi)
		p[i] = pi
	}

	var ans int
	for c := n; c != 1; {
		ans++
		c = p[c]
	}

	fmt.Println(ans)
}

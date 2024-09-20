package main

import "fmt"

func main() {
	var n, x int
	fmt.Scan(&n, &x)

	var ans int
	for i := 0; i < n; i++ {
		var s int
		fmt.Scan(&s)
		if s <= x {
			ans += s
		}
	}

	fmt.Println(ans)
}

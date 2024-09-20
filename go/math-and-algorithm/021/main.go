package main

import "fmt"

func main() {
	var n, r int
	fmt.Scan(&n, &r)

	ans := 1

	for i := 2; i <= n; i++ {
		ans *= i
	}

	for i := 2; i <= r; i++ {
		ans /= i
	}

	for i := 2; i <= n-r; i++ {
		ans /= i
	}

	fmt.Println(ans)
}

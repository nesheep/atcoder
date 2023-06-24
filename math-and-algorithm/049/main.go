package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n+1)
	a[1], a[2] = 1, 1
	for i := 3; i <= n; i++ {
		a[i] = (a[i-1] + a[i-2]) % 1000000007
	}

	fmt.Println(a[n])
}

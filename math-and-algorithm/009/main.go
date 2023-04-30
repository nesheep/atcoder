package main

import "fmt"

func main() {
	var n, s int
	fmt.Scan(&n, &s)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	for i := 1; i <= (1<<n)-1; i++ {
		var sum int
		for j := 0; j <= n-1; j++ {
			if i&(1<<j) != 0 {
				sum += a[j]
			}
		}
		if sum == s {
			fmt.Println("Yes")
			return
		}
	}

	fmt.Println("No")
}

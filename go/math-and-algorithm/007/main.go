package main

import "fmt"

func main() {
	var n, x, y int
	fmt.Scan(&n, &x, &y)

	var cnt int
	for i := 1; i <= n; i++ {
		if i%x == 0 || i%y == 0 {
			cnt++
		}
	}

	fmt.Println(cnt)
}

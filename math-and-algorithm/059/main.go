package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var ans int
	switch n % 4 {
	case 1:
		ans = 2
	case 2:
		ans = 4
	case 3:
		ans = 8
	case 0:
		ans = 6
	}

	fmt.Println(ans)
}

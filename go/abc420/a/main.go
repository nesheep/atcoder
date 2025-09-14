package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	ans := (x + y) % 12
	if ans == 0 {
		ans = 12
	}

	fmt.Println(ans)
}

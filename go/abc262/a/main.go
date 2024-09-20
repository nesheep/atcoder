package main

import "fmt"

func main() {
	var y int
	fmt.Scan(&y)

	ans := y - y%4 + 2
	if y%4 > 2 {
		ans += 4
	}

	fmt.Println(ans)
}

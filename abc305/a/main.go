package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	ans := n - n%5
	if n%5 >= 3 {
		ans += 5
	}
	fmt.Println(ans)
}

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	ans := "First"
	if n%4 == 0 {
		ans = "Second"
	}

	fmt.Println(ans)
}

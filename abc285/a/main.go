package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	var ans bool
	if b == a*2 || b == a*2+1 {
		ans = true
	}

	if ans {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

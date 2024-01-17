package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var ans int
	for n&1 == 0 {
		ans++
		n >>= 1
	}

	fmt.Println(ans)
}

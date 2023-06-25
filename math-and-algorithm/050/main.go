package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	mod := 1000000007
	p := a
	ans := 1
	for i := 0; i < 30; i++ {
		if b&(1<<i) != 0 {
			ans = (ans * p) % mod
		}
		p = (p * p) % mod
	}

	fmt.Println(ans)
}

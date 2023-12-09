package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, k int
	fmt.Scan(&a, &b, &k)

	var ans int
	for {
		if a >= b {
			fmt.Println(ans)
			break
		}
		if a*k/k != a {
			a = math.MaxInt
		} else {
			a *= k
		}
		ans++
	}
}

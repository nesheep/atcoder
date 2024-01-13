package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for x := 0; x <= n; x++ {
		for y := 0; y <= n-x; y++ {
			for z := 0; z <= n-x-y; z++ {
				fmt.Println(x, y, z)
			}
		}
	}
}

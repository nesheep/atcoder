package main

import "fmt"

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	var r, c int
	fmt.Scan(&r, &c)

	dx := minInt(c, 16-c)
	dy := minInt(r, 16-r)

	ans := "black"
	if dx < dy {
		if c%2 == 0 {
			ans = "white"
		}
	} else {
		if r%2 == 0 {
			ans = "white"
		}
	}

	fmt.Println(ans)
}

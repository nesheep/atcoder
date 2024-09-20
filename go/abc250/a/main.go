package main

import "fmt"

func main() {
	var h, w, r, c int
	fmt.Scan(&h, &w, &r, &c)

	var ans int
	if r > 1 {
		ans++
	}
	if r < h {
		ans++
	}
	if c > 1 {
		ans++
	}
	if c < w {
		ans++
	}

	fmt.Println(ans)
}

package main

import "fmt"

func main() {
	var h, m int
	fmt.Scan(&h, &m)

	var ansH, ansM int
	if h < 6 || (h >= 10 && h < 16) {
		ansH, ansM = h, m
	} else if h >= 6 && h < 10 {
		ansH = 10
	} else if h >= 16 && h < 20 {
		ansH = 20
	} else if m < 40 {
		ansH, ansM = h, m
	} else if h < 23 {
		ansH = h + 1
	}

	fmt.Printf("%d %d\n", ansH, ansM)
}

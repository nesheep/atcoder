package main

import "fmt"

func main() {
	var r, c int
	fmt.Scan(&r, &c)
	a := [2][2]int{}
	fmt.Scan(&a[0][0], &a[0][1], &a[1][0], &a[1][1])
	fmt.Println(a[r-1][c-1])
}

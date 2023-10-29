package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	d := x - y
	if d >= -2 && d <= 3 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

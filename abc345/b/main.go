package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	if x < 0 {
		fmt.Println(x / 10)
	} else {
		fmt.Println((x + 9) / 10)
	}
}

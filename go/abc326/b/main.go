package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for {
		a := n / 100
		b := (n - a*100) / 10
		c := n - a*100 - b*10
		if a*b == c {
			fmt.Println(n)
			break
		}
		n++
	}
}

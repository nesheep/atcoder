package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var a int
	var a100, a200, a300, a400 int
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		switch a {
		case 100:
			a100++
		case 200:
			a200++
		case 300:
			a300++
		case 400:
			a400++
		default:
		}
	}

	fmt.Println(a100*a400 + a200*a300)
}

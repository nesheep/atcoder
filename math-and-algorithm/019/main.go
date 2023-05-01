package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var a int
	var a1, a2, a3 int
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		switch a {
		case 1:
			a1++
		case 2:
			a2++
		case 3:
			a3++
		default:
		}
	}

	var cnt int
	cnt += (a1 * (a1 - 1)) / 2
	cnt += (a2 * (a2 - 1)) / 2
	cnt += (a3 * (a3 - 1)) / 2

	fmt.Println(cnt)
}

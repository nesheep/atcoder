package main

import "fmt"

func main() {
	var tm, td, y, m, d int
	fmt.Scan(&tm, &td, &y, &m, &d)

	ny, nm, nd := y, m, d+1
	if nd > td {
		nd = 1
		nm++
	}
	if nm > tm {
		nm = 1
		ny++
	}

	fmt.Println(ny, nm, nd)
}

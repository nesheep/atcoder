package main

import "fmt"

func main() {
	var n, l int
	fmt.Scan(&n, &l)

	var ans int
	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		if a >= l {
			ans++
		}
	}

	fmt.Println(ans)
}

package main

import "fmt"

func main() {
	var n, s int
	fmt.Scan(&n, &s)

	var cnt int
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if i+j <= s {
				cnt++
			}
		}
	}

	fmt.Println(cnt)
}

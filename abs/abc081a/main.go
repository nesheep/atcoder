package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	var cnt int
	for _, v := range s {
		if v == '1' {
			cnt++
		}
	}

	fmt.Println(cnt)
}

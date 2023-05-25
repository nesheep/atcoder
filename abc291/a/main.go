package main

import "fmt"

func main() {
	var s []byte
	fmt.Scan(&s)

	var ans int
	for i, v := range s {
		if v >= byte('A') && v <= byte('Z') {
			ans = i + 1
			break
		}
	}

	fmt.Println(ans)
}

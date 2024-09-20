package main

import "fmt"

func main() {
	var s []byte
	fmt.Scan(&s)

	ans := -1
	for i := range s {
		if s[i] == 'a' {
			ans = i + 1
		}
	}

	fmt.Println(ans)
}

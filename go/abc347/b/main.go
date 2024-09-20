package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	m := make(map[string]struct{})
	for i := 1; i <= len(s); i++ {
		for j := 0; j <= len(s)-i; j++ {
			m[s[j:j+i]] = struct{}{}
		}
	}

	fmt.Println(len(m))
}

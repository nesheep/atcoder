package main

import "fmt"

func main() {
	var s []byte
	fmt.Scan(&s)

	ans := "Yes"
	for i := 1; i < len(s); i += 2 {
		if s[i] != '0' {
			ans = "No"
			break
		}
	}

	fmt.Println(ans)
}

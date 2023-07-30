package main

import "fmt"

func main() {
	var s []byte
	fmt.Scan(&s)

	ans := "No"
	if s[1] == (s[0]+2-'A')%7+'A' && s[2] == (s[1]+2-'A')%7+'A' {
		ans = "Yes"
	}

	fmt.Println(ans)
}

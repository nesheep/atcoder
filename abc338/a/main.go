package main

import "fmt"

func main() {
	var s []byte
	fmt.Scan(&s)

	if s[0] >= 'a' {
		fmt.Println("No")
		return
	}
	for i := 1; i < len(s); i++ {
		if s[i] < 'a' {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}

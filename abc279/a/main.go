package main

import "fmt"

func main() {
	var s []byte
	fmt.Scan(&s)

	var ans int
	for _, c := range s {
		if c == 'w' {
			ans++
		}
		ans++
	}

	fmt.Println(ans)
}

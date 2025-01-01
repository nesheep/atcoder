package main

import "fmt"

func main() {
	var s []byte
	fmt.Scan(&s)

	var ans int
	for i := 0; i < len(s); i++ {
		ans++
		if i < len(s)-1 && s[i] == '0' && s[i+1] == '0' {
			i++
		}
	}

	fmt.Println(ans)
}

package main

import "fmt"

func main() {
	var n int
	var s []byte
	fmt.Scan(&n, &s)

	ans := make([]byte, 0, len(s))
	var d int
	for _, c := range s {
		if c == '(' {
			d++
			ans = append(ans, c)
		} else if c == ')' && d > 0 {
			for {
				p := ans[len(ans)-1]
				ans = ans[:len(ans)-1]
				if p == '(' {
					break
				}
			}
			d--
		} else {
			ans = append(ans, c)
		}
	}

	fmt.Println(string(ans))
}

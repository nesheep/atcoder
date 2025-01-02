package main

import "fmt"

func solve(s, t []byte) bool {
	if len(s) > len(t) {
		return solve(t, s)
	}
	if len(s) == len(t) {
		var d int
		for i := 0; i < len(s); i++ {
			if s[i] != t[i] {
				d++
			}
		}
		return d <= 1
	}
	if len(s)+1 == len(t) {
		var f bool
		for i := 0; i < len(s); i++ {
			if !f && s[i] != t[i] {
				f = true
			}
			if f && s[i] != t[i+1] {
				return false
			}
		}
		return true
	}
	return false
}

func main() {
	var k int
	var s, t []byte
	fmt.Scan(&k, &s, &t)

	if solve(s, t) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

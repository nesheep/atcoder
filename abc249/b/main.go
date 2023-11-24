package main

import "fmt"

func main() {
	var s []byte
	fmt.Scan(&s)

	var large, small bool
	d := true
	m := make(map[byte]bool, 52)

	for _, c := range s {
		if m[c] {
			d = false
			break
		}
		m[c] = true
		if c < 'a' {
			large = true
		} else {
			small = true
		}
	}

	if small && large && d {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

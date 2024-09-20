package main

import "fmt"

func main() {
	var s []byte
	fmt.Scan(&s)

	b1, b2, r1, r2, k := -1, -1, -1, -1, -1
	for i := range s {
		switch s[i] {
		case 'B':
			if b1 < 0 {
				b1 = i
			} else {
				b2 = i
			}
		case 'R':
			if r1 < 0 {
				r1 = i
			} else {
				r2 = i
			}
		case 'K':
			k = i
		}
	}

	if b1%2 != b2%2 && r1 < k && k < r2 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

package main

import "fmt"

func main() {
	var s []byte
	fmt.Scan(&s)

	m := make(map[byte]struct{}, 10)
	for _, v := range s {
		m[v] = struct{}{}
	}

	for b := byte('0'); b <= '9'; b++ {
		if _, ok := m[b]; !ok {
			fmt.Printf("%c\n", b)
		}
	}
}

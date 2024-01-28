package main

import "fmt"

func main() {
	var s []byte
	fmt.Scan(&s)

	m := make(map[byte]int, 30)
	for _, c := range s {
		m[c]++
	}

	ans := byte('a')
	for c := byte('a'); c <= byte('z'); c++ {
		if m[c] > m[ans] {
			ans = c
		}
	}
	fmt.Println(string(ans))
}

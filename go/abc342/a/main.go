package main

import "fmt"

func main() {
	var s []byte
	fmt.Scan(&s)

	m := make(map[byte][]int, 2)
	for i, c := range s {
		m[c] = append(m[c], i+1)
	}

	for _, v := range m {
		if len(v) == 1 {
			fmt.Println(v[0])
			return
		}
	}
}

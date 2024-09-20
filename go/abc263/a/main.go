package main

import "fmt"

func main() {
	var a, b, c, d, e int
	fmt.Scan(&a, &b, &c, &d, &e)

	m := make(map[int]int, 5)
	m[a]++
	m[b]++
	m[c]++
	m[d]++
	m[e]++

	ans := "No"
	if len(m) == 2 && (m[a] == 2 || m[a] == 3) {
		ans = "Yes"
	}

	fmt.Println(ans)
}

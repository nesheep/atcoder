package main

import "fmt"

func main() {
	var n, l, r int
	fmt.Scan(&n, &l, &r)
	l--
	r--

	s := make([]any, n)

	for i := 0; i < l; i++ {
		s[i] = i + 1
	}
	for i := l; i <= r; i++ {
		s[i] = r + l - i + 1
	}
	for i := r + 1; i < n; i++ {
		s[i] = i + 1
	}

	fmt.Println(s...)
}

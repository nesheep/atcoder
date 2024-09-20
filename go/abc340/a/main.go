package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var a, b, d int
	fmt.Scan(&a, &b, &d)

	s := make([]int, (b-a)/d+1)
	ss := make([]string, len(s))
	s[0] = a
	ss[0] = strconv.Itoa(a)
	for i := 1; i < len(s); i++ {
		s[i] = s[i-1] + d
		ss[i] = strconv.Itoa(s[i])
	}

	fmt.Println(strings.Join(ss, " "))
}

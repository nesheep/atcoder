package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	fs := make([]string, 0)
	for i := 2; i*i <= n; i++ {
		for n%i == 0 {
			fs = append(fs, strconv.Itoa(i))
			n /= i
		}
	}
	if n != 1 {
		fs = append(fs, strconv.Itoa(n))
	}

	fmt.Println(strings.Join(fs, " "))
}

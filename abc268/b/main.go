package main

import (
	"fmt"
	"strings"
)

func main() {
	var s, t string
	fmt.Scan(&s, &t)

	ans := "No"
	if strings.HasPrefix(t, s) {
		ans = "Yes"
	}

	fmt.Println(ans)
}

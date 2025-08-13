package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	var s string
	fmt.Scan(&n, &s)

	if strings.HasSuffix(s, "tea") {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

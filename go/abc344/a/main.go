package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)
	ss := strings.Split(s, "|")
	fmt.Printf("%s%s\n", ss[0], ss[2])
}

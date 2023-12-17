package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Printf("%s\n", strings.Repeat(strconv.Itoa(n), n))
}

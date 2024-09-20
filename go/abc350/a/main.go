package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var s string
	fmt.Scan(&s)

	n, _ := strconv.Atoi(s[3:])
	if n < 1 || n > 349 || n == 316 {
		fmt.Println("No")
		os.Exit(0)
	}

	fmt.Println("Yes")
}

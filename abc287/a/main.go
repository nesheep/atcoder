package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(r, &n)

	var s string
	var cnt int
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &s)
		if s == "For" {
			cnt++
		}
	}

	if cnt > n/2 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

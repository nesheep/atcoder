package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rd := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(rd, &n, &m)

	var sum int
	for i := 0; i < n; i++ {
		var h int
		fmt.Fscan(rd, &h)
		sum += h
		if sum > m {
			fmt.Println(i)
			os.Exit(0)
		}
	}

	fmt.Println(n)
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rd := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(rd, &n)

	var a0 int
	fmt.Fscan(rd, &a0)

	for i := 1; i < n; i++ {
		var a int
		fmt.Fscan(rd, &a)
		if a != a0 {
			fmt.Println("No")
			os.Exit(0)
		}
	}

	fmt.Println("Yes")
}

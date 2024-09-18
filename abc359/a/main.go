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

	var cnt int
	for i := 0; i < n; i++ {
		var s string
		fmt.Fscan(rd, &s)
		if s == "Takahashi" {
			cnt++
		}
	}

	fmt.Println(cnt)
}

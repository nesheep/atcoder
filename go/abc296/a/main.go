package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	var n int
	var s []byte
	fmt.Fscan(r, &n, &s)

	ans := true
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			ans = false
			break
		}
	}

	if ans {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

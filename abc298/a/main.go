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

	x := byte('x')
	o := byte('o')

	var ans bool
	for _, v := range s {
		if v == x {
			ans = false
			break
		}
		if v == o {
			ans = true
		}
	}

	if ans {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

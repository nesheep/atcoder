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

	var word string
	var ans bool
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &word)
		if word == "and" || word == "not" || word == "that" || word == "the" || word == "you" {
			ans = true
			break
		}
	}

	if ans {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

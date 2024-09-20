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

	bar := byte('|')
	star := byte('*')

	var f, ans bool
	for _, v := range s {
		if v == bar {
			if f {
				break
			} else {
				f = true
			}
		} else if v == star {
			if f {
				ans = true
			}
			break
		}
	}

	if ans {
		fmt.Println("in")
	} else {
		fmt.Println("out")
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	bt = byte('T')
	ba = byte('A')
)

func main() {
	r := bufio.NewReader(os.Stdin)

	var n int
	var s []byte
	fmt.Fscan(r, &n, &s)

	var ct, ca int
	for _, v := range s {
		switch v {
		case bt:
			ct++
		case ba:
			ca++
		}
	}

	if ct > ca {
		fmt.Println("T")
	} else if ct < ca {
		fmt.Println("A")
	} else {
		if s[n-1] == ba {
			fmt.Println("T")
		} else {
			fmt.Println("A")
		}
	}
}

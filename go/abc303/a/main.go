package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var n int
	var s, t []byte

	fmt.Fscan(rd, &n, &s, &t)

	ans := true
	for i := 0; i < n; i++ {
		c1 := s[i] == t[i]
		c2 := s[i] == '1' && t[i] == 'l'
		c3 := s[i] == 'l' && t[i] == '1'
		c4 := s[i] == '0' && t[i] == 'o'
		c5 := s[i] == 'o' && t[i] == '0'
		if !(c1 || c2 || c3 || c4 || c5) {
			ans = false
			break
		}
	}

	if ans {
		fmt.Fprintln(wr, "Yes")
	} else {
		fmt.Fprintln(wr, "No")
	}
}

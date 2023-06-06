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
	var s []byte
	fmt.Fscan(rd, &n, &s)

	var f bool
	for i := range s {
		switch s[i] {
		case '"':
			f = !f
		case ',':
			if !f {
				s[i] = '.'
			}
		}
	}

	fmt.Fprintf(wr, "%s\n", s)
}

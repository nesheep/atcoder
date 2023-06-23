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

	var h, w int
	fmt.Fscan(rd, &h, &w)

	s := make([][]byte, h)
	for i := range s {
		fmt.Fscan(rd, &s[i])
	}

	for i := range s {
		for j := 0; j < w-1; j++ {
			if s[i][j] == 'T' && s[i][j+1] == 'T' {
				s[i][j] = 'P'
				s[i][j+1] = 'C'
				j++
			}
		}
		fmt.Fprintln(wr, string(s[i]))
	}
}

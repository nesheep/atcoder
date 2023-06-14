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

	var s []byte
	for i := 0; i < 8; i++ {
		fmt.Fscan(rd, &s)
		for j := 0; j < 8; j++ {
			if s[j] == '*' {
				fmt.Fprintf(wr, "%c%d\n", 'a'+byte(j), 8-i)
				return
			}
		}
	}
}

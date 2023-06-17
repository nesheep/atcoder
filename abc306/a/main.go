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
	for _, c := range s {
		fmt.Fprintf(wr, "%c%c", c, c)
	}
	fmt.Fprintln(wr)
}

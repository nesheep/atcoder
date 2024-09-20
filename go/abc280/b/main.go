package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var n int
	fmt.Fscan(rd, &n)

	ss := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &ss[i])
	}

	as := make([]string, n)
	as[0] = strconv.Itoa(ss[0])
	for i := 1; i < n; i++ {
		as[i] = strconv.Itoa(ss[i] - ss[i-1])
	}

	fmt.Fprintln(wr, strings.Join(as, " "))
}

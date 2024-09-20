package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n, k int
	fmt.Fscan(r, &n, &k)

	ss := make([]string, k)
	for i := 0; i < k; i++ {
		fmt.Fscan(r, &ss[i])
	}

	sort.Strings(ss)

	for i := 0; i < k; i++ {
		fmt.Fprintln(w, ss[i])
	}
}

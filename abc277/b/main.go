package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func judge(s string) bool {
	var f bool
	for _, v := range []byte{'H', 'D', 'C', 'S'} {
		if s[0] == v {
			f = true
			break
		}
	}
	if !f {
		return false
	}

	f = false
	for _, v := range []byte{'A', '2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K'} {
		if s[1] == v {
			f = true
			break
		}
	}

	return f
}

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var n int
	fmt.Fscan(rd, &n)

	s := make([]string, n)
	for i := range s {
		fmt.Fscan(rd, &s[i])
	}

	sort.Strings(s)

	ans := "Yes"
	for i := range s {
		if !judge(s[i]) || (i > 0 && s[i-1] == s[i]) {
			ans = "No"
			break
		}
	}

	fmt.Fprintln(wr, ans)
}

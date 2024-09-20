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

	var s, t []byte
	fmt.Fscan(rd, &s, &t)

	ms := make(map[byte]int, 27)
	mt := make(map[byte]int, 27)
	for i := range s {
		ms[s[i]]++
		mt[t[i]]++
	}

	ans := "Yes"
	for c := byte('a'); c <= 'z'; c++ {
		if c == 'a' || c == 't' || c == 'c' || c == 'o' || c == 'd' || c == 'e' || c == 'r' {
			if ms[c] < mt[c] {
				ms['@'] -= mt[c] - ms[c]
			} else if ms[c] > mt[c] {
				mt['@'] -= ms[c] - mt[c]
			}
		} else if ms[c] != mt[c] {
			ans = "No"
			break
		}
	}

	if ms['@'] < 0 || mt['@'] < 0 {
		ans = "No"
	}

	fmt.Fprintln(wr, ans)
}

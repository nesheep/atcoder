package main

import (
	"bytes"
	"fmt"
)

func calc(s []byte) float64 {
	if len(s) < 3 {
		return 0
	}
	if s[0] != 't' || s[len(s)-1] != 't' {
		return 0
	}
	return float64(bytes.Count(s, []byte("t"))-2) / float64(len(s)-2)
}

func main() {
	var s []byte
	fmt.Scan(&s)

	var ans float64
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			if j-i < 3 {
				continue
			}
			a := calc(s[i:j])
			if a > ans {
				ans = a
			}
		}
	}

	fmt.Printf("%.16f\n", ans)
}

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	d := []int{}
	nd := []int{}
	for i := 1; i <= 9; i++ {
		if n%i == 0 {
			d = append(d, i)
			nd = append(nd, n/i)
		}
	}

	s := make([]byte, 0, n+1)
	for i := 0; i <= n; i++ {
		var ok bool
		for j := 0; j < len(d); j++ {
			if i%nd[j] == 0 {
				s = append(s, '0'+byte(d[j]))
				ok = true
				break
			}
		}
		if !ok {
			s = append(s, '-')
		}
	}

	fmt.Println(string(s))
}

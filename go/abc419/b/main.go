package main

import "fmt"

func main() {
	var q int
	fmt.Scan(&q)

	m := make(map[int]int)
	ans := make([]int, 0)

	for i := 0; i < q; i++ {
		var t int
		fmt.Scan(&t)

		switch t {
		case 1:
			var x int
			fmt.Scan(&x)
			m[x]++
		case 2:
			x := 1000
			for k, v := range m {
				if k < x && v > 0 {
					x = k
				}
			}
			ans = append(ans, x)
			m[x]--
		}
	}

	for _, a := range ans {
		fmt.Println(a)
	}
}

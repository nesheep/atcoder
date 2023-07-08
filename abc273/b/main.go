package main

import "fmt"

func main() {
	var x, k int
	fmt.Scan(&x, &k)

	d := 1
	for i := 0; i < k; i++ {
		d *= 10
		if x%d >= 5*d/10 {
			x = x/d*d + d
		} else {
			x = x / d * d
		}
	}

	fmt.Println(x)
}

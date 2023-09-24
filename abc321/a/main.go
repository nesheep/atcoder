package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	ans := "Yes"
	r := n % 10
	n /= 10
	for n > 0 {
		if n%10 <= r {
			ans = "No"
			break
		}
		r = n % 10
		n /= 10
	}

	fmt.Println(ans)
}

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var ans int
	for i := 1; i <= n; i++ {
		var d int
		fmt.Scan(&d)
		if i < 10 {
			if d >= i {
				ans++
			}
			if d >= i*10+i {
				ans++
			}
		} else {
			j := i % 10
			if i/10 != j {
				continue
			}
			if d >= j {
				ans++
			}
			if d >= j*10+j {
				ans++
			}
		}
	}

	fmt.Println(ans)
}

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	var cnt int
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				for l := k + 1; l < n; l++ {
					for m := l + 1; m < n; m++ {
						if a[i]+a[j]+a[k]+a[l]+a[m] == 1000 {
							cnt++
						}
					}
				}
			}
		}
	}

	fmt.Println(cnt)
}

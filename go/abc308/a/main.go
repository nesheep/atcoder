package main

import "fmt"

func main() {
	ans := "Yes"
	s := [9]int{100}
	for i := 1; i <= 8; i++ {
		fmt.Scan(&s[i])
		c1 := s[i] >= s[i-1]
		c2 := s[i] >= 100 && s[i] <= 675
		c3 := s[i]%25 == 0
		if !c1 || !c2 || !c3 {
			ans = "No"
			break
		}
	}
	fmt.Println(ans)
}

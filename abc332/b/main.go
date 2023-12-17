package main

import "fmt"

func main() {
	var k, g, m int
	fmt.Scan(&k, &g, &m)

	var ag, am int
	for i := 0; i < k; i++ {
		if ag == g {
			ag = 0
		} else if am == 0 {
			am = m
		} else if am <= g-ag {
			ag += am
			am = 0
		} else {
			am -= g - ag
			ag = g
		}
	}

	fmt.Println(ag, am)
}

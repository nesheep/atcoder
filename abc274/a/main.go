package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	if a == b {
		fmt.Println("1.000")
		return
	}

	c := b * 10000 / a
	c1 := c / 1000
	c2 := c % 1000 / 100
	c3 := c % 100 / 10
	c4 := c % 10
	if c4 >= 5 {
		c3++
	}

	fmt.Printf("0.%d%d%d\n", c1, c2, c3)
}

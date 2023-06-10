package main

import "fmt"

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	var p, q string
	fmt.Scan(&p, &q)

	sum := make(map[string]int, 6)
	sum["B"] = sum["A"] + 3
	sum["C"] = sum["B"] + 1
	sum["D"] = sum["C"] + 4
	sum["E"] = sum["D"] + 1
	sum["F"] = sum["E"] + 5
	sum["G"] = sum["F"] + 9

	fmt.Println(abs(sum[p] - sum[q]))
}

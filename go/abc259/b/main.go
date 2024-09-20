package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, d float64
	fmt.Scan(&a, &b, &d)

	d = d * math.Pi / 180
	fmt.Printf("%.20f %.20f\n", a*math.Cos(d)-b*math.Sin(d), a*math.Sin(d)+b*math.Cos(d))
}

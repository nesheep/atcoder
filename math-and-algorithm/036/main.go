package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	var a, b, h, m float64
	fmt.Fscan(r, &a, &b, &h, &m)

	// 時計の中心を原点、12時方向をx軸正、3時方向をy軸正とする
	// c1はx軸と時針のなす角
	// c2はx軸と分針のなす角
	// 時針の先端 P1=(x1,y1)
	// 分針の先端 P2=(x2,y2)
	c1 := math.Pi/6*h + math.Pi/360*m
	c2 := math.Pi / 30 * m
	x1 := a * math.Cos(c1)
	y1 := a * math.Sin(c1)
	x2 := b * math.Cos(c2)
	y2 := b * math.Sin(c2)

	// P1とP2の距離を求める
	ans := math.Sqrt(math.Pow(x1-x2, 2) + math.Pow(y1-y2, 2))

	fmt.Printf("%.20f\n", ans)
}

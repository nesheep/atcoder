package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type pattern int

const (
	p1 pattern = iota + 1 // 角ABCが90度より大きい
	p2                    // 角ABC,ACBがともに90度以下
	p3                    // 角ACBが90度より大きい
)

func main() {
	r := bufio.NewReader(os.Stdin)

	var ax, ay, bx, by, cx, cy float64
	fmt.Fscan(r, &ax, &ay, &bx, &by, &cx, &cy)

	// ベクトルBA,BC,CA,CBの各成分を求める
	bax := ax - bx
	bay := ay - by
	bcx := cx - bx
	bcy := cy - by
	cax := ax - cx
	cay := ay - cy
	cbx := bx - cx
	cby := by - cy

	// どのパターンに当てはまるか判定
	p := p2
	if bax*bcx+bay*bcy < 0 {
		p = p1
	}
	if cax*cbx+cay*cby < 0 {
		p = p3
	}

	// 点と線分の距離を求める
	var ans float64
	switch p {
	case p1:
		ans = math.Sqrt(math.Pow(bax, 2) + math.Pow(bay, 2))
	case p2:
		s := math.Abs(bax*cay - bay*cax)
		bc := math.Sqrt(math.Pow(bcx, 2) + math.Pow(bcy, 2))
		ans = s / bc
	case p3:
		ans = math.Sqrt(math.Pow(cax, 2) + math.Pow(cay, 2))
	default:
	}

	fmt.Printf("%.12f\n", ans)
}

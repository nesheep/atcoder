package main

import (
	"bufio"
	"fmt"
	"os"
)

type pair struct {
	x, y int
}

func vector(a, b pair) pair {
	return pair{
		x: b.x - a.x,
		y: b.y - a.y,
	}
}

func cross(a, b pair) int {
	return a.x*b.y - a.y*b.x
}

func swap(a, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	r := bufio.NewReader(os.Stdin)

	var a, b, c, d pair
	fmt.Fscan(r, &a.x, &a.y)
	fmt.Fscan(r, &b.x, &b.y)
	fmt.Fscan(r, &c.x, &c.y)
	fmt.Fscan(r, &d.x, &d.y)

	ab := vector(a, b)
	ac := vector(a, c)
	ad := vector(a, d)
	crossBAC := cross(ab, ac)
	crossBAD := cross(ab, ad)

	cd := vector(c, d)
	ca := vector(c, a)
	cb := vector(c, b)
	crossDCA := cross(cd, ca)
	crossDCB := cross(cd, cb)

	if crossBAC == 0 && crossBAD == 0 {
		if a.x == b.x {
			ay := a.y
			by := b.y
			cy := c.y
			dy := d.y
			if ay > by {
				swap(&ay, &by)
			}
			if cy > dy {
				swap(&cy, &dy)
			}

			if max(ay, cy) <= min(by, dy) {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}
			return
		}

		ax := a.x
		bx := b.x
		cx := c.x
		dx := d.x
		if ax > bx {
			swap(&ax, &bx)
		}
		if cx > dx {
			swap(&cx, &dx)
		}

		if max(ax, cx) <= min(bx, dx) {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
		return
	}

	var cond1, cond2 bool
	if crossBAC >= 0 && crossBAD <= 0 || crossBAC <= 0 && crossBAD >= 0 {
		cond1 = true
	}
	if crossDCA >= 0 && crossDCB <= 0 || crossDCA <= 0 && crossDCB >= 0 {
		cond2 = true
	}

	if cond1 && cond2 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

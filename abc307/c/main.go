package main

import (
	"bufio"
	"fmt"
	"os"
)

func scanSheet(rd *bufio.Reader) [][]byte {
	var h, w int
	fmt.Fscan(rd, &h, &w)
	s := make([][]byte, h)
	for i := range s {
		fmt.Fscan(rd, &s[i])
	}
	return s
}

func getRange(a [][]byte) [4]int {
	top, bottom, left, right := -1, -1, -1, -1
	for i := range a {
		for j := range a[i] {
			if a[i][j] == '#' {
				if top == -1 {
					top = i
				}
				bottom = i
				if left == -1 || left > j {
					left = j
				}
				if right < j {
					right = j
				}
			}
		}
	}
	return [4]int{top, bottom, left, right}
}

func cutSheet(a [][]byte) [][]byte {
	r := getRange(a)
	s := make([][]byte, r[1]-r[0]+1)
	for i := range s {
		s[i] = make([]byte, r[3]-r[2]+1)
		copy(s[i], a[r[0]+i][r[2]:r[3]+1])
	}
	return s
}

func toBoolSlice(a [][]byte) [][]bool {
	b := make([][]bool, len(a))
	for i := range b {
		b[i] = make([]bool, len(a[i]))
		for j := range b[i] {
			if a[i][j] == '#' {
				b[i][j] = true
			}
		}
	}
	return b
}

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	a := scanSheet(rd)
	b := scanSheet(rd)
	x := scanSheet(rd)

	ab := toBoolSlice(cutSheet(a))
	bb := toBoolSlice(cutSheet(b))
	xb := toBoolSlice(cutSheet(x))

	ans := "No"

	if len(xb) < len(ab) || len(xb[0]) < len(ab[0]) || len(xb) < len(bb) || len(xb[0]) < len(bb[0]) {
		fmt.Fprintln(wr, ans)
		return
	}

All:
	for ai := 0; ai <= len(xb)-len(ab); ai++ {
		for aj := 0; aj <= len(xb[0])-len(ab[0]); aj++ {
			for bi := 0; bi <= len(xb)-len(bb); bi++ {
				for bj := 0; bj <= len(xb[0])-len(bb[0]); bj++ {
					ok := true
				One:
					for i := 0; i < len(xb); i++ {
						for j := 0; j < len(xb[0]); j++ {
							ca := i >= ai && i < ai+len(ab) && j >= aj && j < aj+len(ab[0])
							cb := i >= bi && i < bi+len(bb) && j >= bj && j < bj+len(bb[0])
							if xb[i][j] != (ca && ab[i-ai][j-aj] || cb && bb[i-bi][j-bj]) {
								ok = false
								break One
							}
						}
					}
					if ok {
						ans = "Yes"
						break All
					}
				}
			}
		}
	}

	fmt.Fprintln(wr, ans)
}

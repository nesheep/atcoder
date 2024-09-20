package main

import "fmt"

const mod = 998244353

func main() {
	var n, m uint64
	fmt.Scan(&n, &m)

	var ans uint64
	for i := uint64(0); i < 60; i++ {
		if 1<<i&m != 0 {
			c := (n + 1) / (2 << i)
			r := (n + 1) % (2 << i)
			ans += (c << i) % mod
			if r > 1<<i {
				ans += (r - (1 << i)) % mod
			}
		}
	}

	fmt.Println(ans % mod)
}

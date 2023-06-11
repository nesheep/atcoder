package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"sort"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var h, w int
	fmt.Fscan(rd, &h, &w)

	ss := make([][]byte, w)
	for i := 0; i < w; i++ {
		ss[i] = make([]byte, h)
	}

	ts := make([][]byte, w)
	for i := 0; i < w; i++ {
		ts[i] = make([]byte, h)
	}

	var tmp []byte
	for i := 0; i < h; i++ {
		fmt.Fscan(rd, &tmp)
		for j := range tmp {
			ss[j][i] = tmp[j]
		}
	}
	for i := 0; i < h; i++ {
		fmt.Fscan(rd, &tmp)
		for j := range tmp {
			ts[j][i] = tmp[j]
		}
	}

	sort.Slice(ss, func(i, j int) bool { return string(ss[i]) < string(ss[j]) })
	sort.Slice(ts, func(i, j int) bool { return string(ts[i]) < string(ts[j]) })

	if reflect.DeepEqual(ss, ts) {
		fmt.Fprintln(wr, "Yes")
	} else {
		fmt.Fprintln(wr, "No")
	}
}

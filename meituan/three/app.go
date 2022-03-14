package main

import (
	"fmt"
	"sort"
)

func main() {

	var n, m int
	var segs [][2]int
	fmt.Scan(&n, &m)
	segs = make([][2]int, 0)
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		segs = append(segs, [2]int{a, b})
	}
	sort.Slice(segs, func(i, j int) bool {
		if segs[i][0] != segs[j][0] {
			return segs[i][0] < segs[j][0]
		} else {
			return segs[i][1] < segs[j][1]
		}
	})

	res := 1
	l := segs[n-1][0]

	for i := n - 2; i >= 0; i-- {
		if segs[i][1] >= l {
			continue
		} else {
			l = segs[i][0]
			res++
		}
	}

	fmt.Println(res)
}

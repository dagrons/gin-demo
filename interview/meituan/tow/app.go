package main

import "fmt"

// neg[i]: 第i个-1的位置

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	f := make([]int, n)
	g := make([]int, n)

	if a[0] == -1 {
		f[0] = 1
	} else {
		g[0] = 1
	}

	for i := 1; i < n; i++ {
		if a[i] == -1 {
			f[i] = g[i-1] + 1
			g[i] = f[i-1]
		} else {
			f[i] = f[i-1]
			g[i] = g[i-1] + 1
		}
	}

	res := 0
	for i := 0; i < n; i++ {
		res += g[i]
	}
	fmt.Println(res)
}

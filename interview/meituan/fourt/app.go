package main

import "fmt"

// good[i]: 对于时刻i，房间good[i]能在最长时间内不受到伤害

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	a := make([][]bool, n)
	for i := 0; i < n; i++ {
		a[i] = make([]bool, m)
	}
	for i := 0; i < m; i++ {
		var x int
		fmt.Scan(&x)
		a[x-1][i] = true
	}

	good := make([]int, m)
	goodCnt := make([]int, m)
	for i := 0; i < m; i++ {
		cns := make([]int, n)
		for j := 0; j < n; j++ {
			k := i
			for k < m && !a[j][k] {
				k++
			}
			cns[j] = k
		}

		maxj := 0
		maxcns := 0
		for p, v := range cns {
			if v > maxcns {
				maxj = p
				maxcns = v
			}
		}
		good[i] = maxj
		goodCnt[i] = maxcns
	}

	res := 0
	i := 0
	for i < m {
		i = goodCnt[i]
		res++
	}
	fmt.Println(res)
}

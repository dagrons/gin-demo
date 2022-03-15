package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	bd := make([][]byte, 8)
	for i := 0; i < 8; i++ {
		var s string
		fmt.Scan(&s)
		bd[i] = []byte(s)
	}
	ss := make([][]byte, 8)
	for i := 0; i < 8; i++ {
		var s string
		fmt.Scan(&s)
		ss[i] = []byte(s)
	}
	for i := 0; i < n; i++ {
		var x, y int
		var op string
		fmt.Scan(&x, &y, &op)
		x--
		y--
		color := bd[x][y]
		cnt := 0
		dfs(bd, x, y, color, &cnt)
		fmt.Print(cnt)
		// move
		for r := 0; r < 8; r++ {
			for c := 0; c < 8; c++ {
				for bd[r][c] == '*' {
					fmt.Println()
					switch op {
					case "w":
						for k := r; k < 7; k++ {
							bd[k][c] = bd[k+1][c]
						}
						bd[7][c] = ss[7-c][0]
						ss[7-c] = ss[7-c][1:]
					case "a":
						for k := c; k < 7; k++ {
							bd[r][k] = bd[r][k+1]
						}
						bd[r][7] = ss[r][0]
						ss[r] = ss[r][1:]
					case "s":
						for k := r; k >= 1; k-- {
							bd[k][c] = bd[k-1][c]
						}
						bd[0][c] = ss[c][0]
						ss[c] = ss[c][1:]
					case "d":
						for k := c; k >= 1; k-- {
							bd[r][k] = bd[r][k-1]
						}
						bd[7-r][0] = ss[7-r][0]
						ss[7-r] = ss[7-r][1:]
					}
				}
			}
		}
	}
}

func dfs(bd [][]byte, x, y int, color byte, cnt *int) {
	if x < 0 || y < 0 || x >= 8 || y >= 8 || bd[x][y] != color {
		return
	}
	bd[x][y] = '*'
	*cnt++
	dx := []int{0, 1, 0, -1}
	dy := []int{1, 0, -1, 0}
	for i := 0; i < 4; i++ {
		tx := x + dx[i]
		ty := y + dy[i]
		dfs(bd, tx, ty, color, cnt)
	}
}

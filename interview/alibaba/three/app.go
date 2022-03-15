package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	// load board
	board := make([][]byte, 8)
	for i := 0; i < 8; i++ {
		board[i] = make([]byte, 8)
		var s string
		fmt.Scan(&s)
		for j := 0; j < 8; j++ {
			board[i][j] = s[j]
		}
	}

	// load boardRepo
	boardRepo := make([][]byte, 8)
	for i := 0; i < 8; i++ {
		var tss string
		fmt.Scan(&tss)
		boardRepo[i] = make([]byte, len(tss))
		for j := 0; j < len(tss); j++ {
			boardRepo[i][j] = tss[j]
		}
	}

	// 开始操作
	for i := 0; i < n; i++ {
		var x, y int
		var op string
		fmt.Scan(&x, &y, &op)
		x--
		y--
		// 在x，y处进行消除，将所有联通块同色消除为*
		cl := board[x][y]
		cnt := 0
		dfs(board, cl, x, y, &cnt)
		fmt.Println(cnt)
		move(board, op)
		fix(board, op, boardRepo)
	}
}

func move(board [][]byte, op string) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == '*' {
				switch op {
				case "w":
					k := 0
					for i+k < 8 && board[i+k][j] == '*' {
						k++
					}
					for r := i; r < len(board); r++ {
						if i < 0 || j < 0 || i >= 8 || j >= 8 {
							board[r][j] = '*'
						} else {
							board[r][j] = board[r+k][j]
						}
					}
				case "a":
					k := 0
					for j+k < 8 && board[i][j+k] == '*' {
						k++
					}
					for c := 0; c < len(board[0]); c++ {
						if i < 0 || j < 0 || i >= 8 || j >= 8 {
							board[i][c] = '*'
						} else {
							board[i][c] = board[i][c+k]
						}
					}
				case "s":
					k := 0
					for i-k >= 0 && board[i-k][j] == '*' {
						k++
					}
					for r := 7; r >= 0; r-- {
						if i < 0 || j < 0 || i >= 8 || j >= 8 {
							board[r][j] = '*'
						} else {
							board[r][j] = board[r-k][j]
						}
					}
				case "d":
					k := 0
					for j-k >= 0 && board[i][j-k] == '*' {
						k++
					}
					for c := 7; c >= 0; c-- {
						if i < 0 || j < 0 || i >= 8 || j >= 8 {
							board[i][c] = '*'
						} else {
							board[i][c] = board[i][c-k]
						}
					}
				}
			}
		}
	}
}

func fix(board [][]byte, op string, boardRepo [][]byte) {

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if board[i][j] == '*' {
				switch op {
				case "w":
					t := boardRepo[7-i][0]
					boardRepo[7-i] = boardRepo[7-i][1:] // pop
					board[i][j] = t
				case "a":
					t := boardRepo[j][0]
					boardRepo[j] = boardRepo[j][1:]
					board[i][j] = t
				case "s":
					t := boardRepo[i][0]
					boardRepo[i] = boardRepo[i][1:]
					board[i][j] = t
				case "d":
					t := boardRepo[7-j][0]
					boardRepo[7-j] = boardRepo[7-j][1:]
					board[i][j] = t
				}
			}
		}
	}
}

// 从(x, y)开始消除所有联通块同色
func dfs(board [][]byte, cl byte, x, y int, cnt *int) {
	if board[x][y] != cl {
		return
	}
	board[x][y] = '*'
	*cnt++
	dx := []int{0, 1, 0, -1}
	dy := []int{1, 0, -1, 0}
	for i := 0; i < 4; i++ {
		tx := x + dx[i]
		ty := y + dy[i]
		if tx < 0 || ty < 0 || tx >= 8 || ty >= 8 {
			continue
		}
		dfs(board, cl, tx, ty, cnt)
	}
}

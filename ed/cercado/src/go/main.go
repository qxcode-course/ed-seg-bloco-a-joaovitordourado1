package main

import (
	"bufio"
	"fmt"
	"os"
)

func dfs(board [][]byte, r, c int) {
	rows := len(board)
	columns := len(board[0])
	if r <0 || r >= rows || c < 0 || c>= columns|| board[r][c] != 'O' {
		return
	}
	board[r][c]= 'K'
	
	dfs(board, r+1,c)
	dfs(board, r-1,c)
	dfs(board, r,c+1)
	dfs(board, r,c-1)


	
}

// NÃO ALTERE A ASSINATURA DA FUNÇÃO solve
func solve(board [][]byte) {
	if len(board) == 0||len(board[0]) == 0{
		return
	}
	rows := len(board)
	columns := len(board[0])

	for r := 0;r < rows; r++{
		dfs(board, r, 0)
        dfs(board, r, columns-1)
	}
	for c:= 0; c< columns; c++{
		dfs(board, 0, c)
		dfs(board, rows-1, c-1)
	}
	for r:=0; r < rows; r++{
		for c:= 0; c  < columns; c++{
			if board[r][c] == 'O'{
				board[r][c] ='X'
			}else if board[r][c]== 'K'{
				board[r][c] ='O'
			}
		}
	}
}

// NÃO ALTERE A MAIN
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var nrows, ncols int
	fmt.Sscanf(scanner.Text(), "%d %d", &nrows, &ncols)
	board := make([][]byte, nrows)
	for i := 0; i < nrows; i++ {
		scanner.Scan()
		board[i] = []byte(scanner.Text())
	}
	solve(board)
	for _, row := range board {
		fmt.Println(string(row))
	}
}

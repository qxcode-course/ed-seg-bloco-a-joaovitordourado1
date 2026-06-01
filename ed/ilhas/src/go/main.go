package main

import (
	"bufio"
	"fmt"
	"os"
)

// Não modifique a assinatura da função numIslands
// Ela é a função que será chamada no LeetCode para resolver o problema
func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	n := 0
	r := len(grid)
	c := len(grid[0])
	for i := 0; i < r; i++ {
		for k := 0; k < c; k++ {
			if grid[i][k] == '1' {
				n++
				backtrack(grid, i, k)
			}
		}
	}
	return n
}
func backtrack(l[][]byte, r,c int ){
	if r < 0 || c < 0 || r >= len(l) || c >= len(l[0]) || l[r][c] == '0' {
		return
	}
	l[r][c] = '0'
	backtrack(l, r-1, c)  
	backtrack(l, r+1, c) 
	backtrack(l, r, c-1) 
	backtrack(l, r, c+1) 
}

// Não modifique a função main
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc int
	fmt.Sscanf(line, "%d %d", &nl, &nc)
	grid := make([][]byte, nl)
	for i := 0; i < nl; i++ {
		scanner.Scan()
		grid[i] = []byte(scanner.Text())
	}
	result := numIslands(grid)
	fmt.Println(result)
}

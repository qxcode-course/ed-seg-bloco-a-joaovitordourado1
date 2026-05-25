package main

import (
	"bufio"
	"fmt"
	"os"
)

// Não mude a assinatura desta função, ela é a função chamada pelo LeetCode
func exist(grid [][]byte, word string) bool {
	_, _ = grid, word
	rM,cM := len(grid), len(grid[0])
	var backtrack func(r, c, i int) bool
	backtrack=func(r,c, i int) bool{
		if i == len(word){
			return true
		}
		if r < 0 || r >= rM || c < 0 || c >= cM ||grid[r][c]!= word[i]{
            return false
			
        }
		t := grid[r][c]
		grid[r][c] = '#'
		res :=backtrack(r+1,c,i+1)||
		backtrack(r-1,c,i+1)||
		backtrack(r,c+1,i+1)||
		backtrack(r,c-1,i+1)
		grid[r][c] = t
		return res

	}
	for r := 0; r < rM; r++ {
        for c := 0; c < cM; c++ {
            if backtrack(r, c,0) {
                return true
            }
        }
    }
    return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var word string
	fmt.Sscanf(scanner.Text(), "%s", &word)
	grid := make([][]byte, 0)
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}
	if exist(grid, word) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}

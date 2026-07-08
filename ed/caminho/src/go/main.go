package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	l, c int
}

func (p Pos) getNeig() []Pos {
	return []Pos {
		{l:p.l-1,c:p.c},
		{l:p.l+1,c: p.c},
		{l:p.l,c:p.c-1},
		{l:p.l,c:p.c+1},
	}
}

func inside(grid [][]rune, pos Pos) bool {
	nrows := len(grid)
	ncols := len(grid[0])
	return pos.l >= 0 && pos.l < nrows && pos.c >= 0 && pos.c < ncols
}

func match(grid [][]rune, pos Pos, char rune) bool {
	return inside(grid, pos) && grid[pos.l][pos.c] == char
}

func search(grid [][]rune, startPos Pos, endPos Pos) {
	queue := NewQueue[Pos]()
	queue.Enqueue(startPos)

	visitados := make(map[Pos]bool)
	visitados[startPos] = true


	indopor := make (map[Pos]Pos)
	end:= false

	for !queue.IsEmpty(){
		current, _:= queue.Dequeue()

		if current == endPos{
			end = true
			break
		}
		for _, vizinho := range current.getNeig(){
	
	
	
		if  inside(grid,vizinho) && grid[vizinho.l][vizinho.c]!='#' && !visitados[vizinho]{
			visitados[vizinho] = true
				indopor[vizinho] = current
				queue.Enqueue(vizinho)
		}
	
	
	}

	}
	if end {
		x := endPos
		for{
			grid [x.l][x.c] = '.'
			if x == startPos{
				break
			}
			x = indopor[x]
		}
	}
}



func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var nl, nc int
	scanner.Scan()
	line := scanner.Text()
	fmt.Sscanf(line, "%d %d", &nl, &nc)
	mat := make([][]rune, nl) // Inicializa a matriz de runes

	// Carregando matriz
	for i := range nl {
		scanner.Scan()
		line := scanner.Text()
		mat[i] = []rune(line)
	}

	var inicio, fim Pos

	// Procurando inicio e fim e colocando ' ' nas posições iniciais
	for l := range nl {
		for c := range nc {
			if mat[l][c] == 'I' {
				mat[l][c] = ' '
				inicio = Pos{l, c}
			}
			if mat[l][c] == 'F' {
				mat[l][c] = ' '
				fim = Pos{l, c}
			}
		}
	}

	search(mat, inicio, fim)

	for _, line := range mat {
		fmt.Println(string(line)) // Converte o slice de runes de volta para string para imprimir
	}
}

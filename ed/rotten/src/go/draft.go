package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


type cor struct{
    l int
    c int
}

func orangesRotting(grid [][]int) int {
    li := len(grid)
    co:= len(grid[0])
    

    qu:= []cor{}
    fr :=0

    for i := 0; i < li; i++ {
        for j := 0; j < co; j++ {
            if grid[i][j]==2{
                qu =append(qu,cor{l:i,c:j})
        }else if grid[i][j] ==1{
            fr++
        }
        }
    }
    minutos := 0
	direcoes := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
    for len(qu) > 0 && fr > 0{
        tamanhoLote := len(qu)
        for i := 0; i < tamanhoLote; i++ {
                ac := qu[0]
                qu=qu [1:]
                for _, dir := range direcoes {
				    novaL := ac.l + dir[0]
				    novaC := ac.c + dir[1]
                    if novaL >= 0 && novaL < li && novaC >= 0 && novaC < co && grid[novaL][novaC] == 1{
                        grid[novaL][novaC] = 2
                        fr --
                        qu = append(qu, cor{l: novaL, c: novaC})
                    }

                
    }
   
    
}
minutos++
    }
    if fr > 0 {
		return -1
	}
     return minutos
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {
		return
	}
	line := scanner.Text()

	var nl, nc int
	fmt.Sscanf(line, "%d %d", &nl, &nc)

	mat := make([][]int, nl)

	for i := 0; i < nl; i++ {
		scanner.Scan()
		linhaAtual := scanner.Text()

		partes := strings.Fields(linhaAtual)
		mat[i] = make([]int, nc)

		for j := 0; j < nc; j++ {
			valor, _ := strconv.Atoi(partes[j])
			mat[i][j] = valor
		}
	}

	resultado := orangesRotting(mat)
	fmt.Println(resultado)
}
// professor tive que usar ia para ajudar a ler as entradas, peguei como ref a questaop anterior
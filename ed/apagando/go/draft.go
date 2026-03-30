package main

import "fmt"

func main() {
	var qtdInicial int
	fmt.Scan(&qtdInicial)
	pessoasFila := make([]int, qtdInicial)
	for i := 0; i < qtdInicial; i++ {
		fmt.Scan(&pessoasFila[i])
	}
	qntSai := 0
	fmt.Scan(&qntSai)
	vouTirar := make(map[int]bool)
	for i := 0; i < qntSai; i++ {
		var idSaindo int
		fmt.Scan(&idSaindo)
		vouTirar[idSaindo] = true
	}
	for i := 0; i < qtdInicial; i++ {
		if !vouTirar[pessoasFila[i]] {
			fmt.Printf("%v ", pessoasFila[i])
		}
	}
	fmt.Println()
}

package main

import "fmt"

func imprimir(pos int, pessoas []int) {
	fmt.Print("[")
	for i := 0; i < len(pessoas); i++ {
		if pessoas[i] != 0 {
			if i == pos {
				fmt.Printf(" %d>", pessoas[i])
			} else {
				fmt.Printf(" %d", pessoas[i])
			}

		}
	}
	fmt.Println(" ]")
}

func matarproximo(numeroPessoas int, quemMata int) {
	pessoas := make([]int, numeroPessoas)
	for i := 0; i < numeroPessoas; i++ {
		pessoas[i] = i + 1
	}
	vivos := numeroPessoas
	pos := quemMata - 1
	for vivos > 0 {
		imprimir(pos, pessoas)
		if vivos == 1 {
			break
		}

		indiceMorto := (pos + 1) % numeroPessoas
		for pessoas[indiceMorto] == 0 {
			indiceMorto = (indiceMorto + 1) % numeroPessoas
		}
		pessoas[indiceMorto] = 0
		vivos--

		pos = (indiceMorto + 1) % numeroPessoas

		for pessoas[pos] == 0 {
			pos = (pos + 1) % numeroPessoas
		}
	}
}

func main() {
	numeroPessoas := 0
	quemMata := 0
	fmt.Scan(&numeroPessoas, &quemMata)
	matarproximo(numeroPessoas, quemMata)
}

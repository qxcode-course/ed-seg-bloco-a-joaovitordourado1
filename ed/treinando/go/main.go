package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)



func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	if len(slice) == 1 {
		return fmt.Sprintf("%d", slice[0])
	}
	return fmt.Sprintf("%d%s", slice[0], sep) + Join(slice[1:], sep)
}
func tostr(vet []int) string {
	if len(vet) == 0 {
		return "[]"
	}
	if len(vet) == 1 {
		return fmt.Sprintf("[%d]", vet[0])
	}

	return "[" +Join(vet,", ")+ "]"
}

func tostrrev(vet []int) string {
	_ = vet
	return ""
}

// reverse: inverte os elementos do slice
func reverse(vet []int) {
	_ = vet
}

// sum: soma dos elementos do slice
func sum(vet []int) int {
	if len(vet)== 1{
		return 0
	}
	return vet[0] + sum(vet[1:])
}


// mult: produto dos elementos do slice
func mult(vet []int) int {
	if len(vet)== 1{
		return 0
	}
	return vet[0] * sum(vet[1:])
}

// min: retorna o índice e valor do menor valor
// crie uma função recursiva interna do modelo
// var rec func(v []int) (int, int)
// para fazer uma recursão que retorna valor e índice
func min(vet []int) int {
	_ = vet
	return 0
}

func main() {
	var vet []int
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Fields(line)
		fmt.Println("$" + line)

		switch args[0] {
		case "end":
			return
		case "read":
			vet = nil
			for _, arg := range args[1:] {
				if val, err := strconv.Atoi(arg); err == nil {
					vet = append(vet, val)
				}
			}
		case "tostr":
			fmt.Println(tostr(vet))
		case "torev":
			fmt.Println(tostrrev(vet))
		case "reverse":
			reverse(vet)
		case "sum":
			fmt.Println(sum(vet))
		case "mult":
			fmt.Println(mult(vet))
		case "min":
			fmt.Println(min(vet))
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}

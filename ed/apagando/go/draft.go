package main
import "fmt"
func main() {
	var qtdInicial int
	fmt.Scan(&qtdInicial)
	pessoasFila := make([]int, qtdInicial)
	for   i:= 0 ; i< qtdInicial ; i++{
		fmt.Scan(&pessoasFila[i])
	}
	qntSai := 0
	fmt.Scan(&qntSai)


}

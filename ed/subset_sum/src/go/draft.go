package main

import "fmt"

func subset(l []int, alvo int, soma int, i int) bool {
	if soma == alvo {
		return true
	}
	if soma > alvo {
		return false
	}
	if i == len(l) {
		return false
	}
	return subset(l, alvo, soma+l[i], i+1) || subset(l, alvo, soma, i+1)
}
func main() {
	n := 0
	alvo := 0
	fmt.Scan(&n, &alvo)
	valores := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&valores[i])
	}
	resultado := subset(valores, alvo, 0, 0)
	fmt.Println(resultado)
}

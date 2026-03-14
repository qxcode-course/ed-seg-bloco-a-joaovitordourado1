package main

import "fmt"

func main() {
	var n, ganhador, aux int
	fmt.Scan(&n)
	ganhador = -1
	aux = 99999
	for i := 0; i < n; i++ {
		var n1, n2, dif int
		fmt.Scan(&n1, &n2)
		if n1 < 10 || n2 < 10 {
			continue
		}
		dif = n1 - n2
		if dif < 0 {
			dif *= -1
		}
		if dif < aux {
			aux = dif
			ganhador = i
		}
	}
	if ganhador == -1 {
		fmt.Print("sem ganhador\n")
	} else {
		fmt.Println(ganhador)
	}
}

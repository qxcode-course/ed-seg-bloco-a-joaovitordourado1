package main

import "fmt"

func main() {
	var n, n1, cont int
	fmt.Scan(&n, &n1)

	figurinhas := make([]int, n1)
	for i := 0; i < n1; i++ {
		fmt.Scan(&figurinhas[i])
	}
	flag := true
	for i := 0; i < n1; i++ {
		for j := 0; j < i; j++ {
			if figurinhas[i] == figurinhas[j] {
				if flag {
					fmt.Printf("%d", figurinhas[i])
					flag = false
				} else {
					fmt.Printf(" %d", figurinhas[i])
				}

				cont++
				break
			}
		}
	}
	if cont == 0 {
		fmt.Print("N")
	}
	fmt.Print("\n")
	flag = true
	cont = 0
	for i := 1; i <= n; i++ {
		j := 0
		for j < n1 {
			if figurinhas[j] == i {
				break
			}
			j++
		}
		if j == n1 {
			if flag {
				fmt.Printf("%d", i)
				flag = false
				cont++
			} else {
				fmt.Printf(" %d", i)
				cont++
			}
		}
	}
	if cont == 0 {
		fmt.Print("N")
	}
	fmt.Print("\n")
}

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	animais := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&animais[i])
	}
	casais := 0
	for i := 0; i < n; i++ {
		if animais[i] == 0 {
			continue
		}

		for j := i + 1; j < n; j++ {
			if animais[j] != 0 && animais[i]+animais[j] == 0 {
				animais[i] = 0
				animais[j] = 0
				casais++
				break
			}
		}
	}
	fmt.Println(casais)
}

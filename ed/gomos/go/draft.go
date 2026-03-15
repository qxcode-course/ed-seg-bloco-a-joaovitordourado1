package main

import "fmt"

func main() {
	var d string
	var n int
	fmt.Scan(&n, &d)
	gomos := make([]int, n*2)
	for i := 0; i < n*2; i++ {
		fmt.Scan(&
            gomos[i])
	}
	for i := n*2 - 1; i > 1; i-- {
		gomos[i] = gomos[i-2]
	}
	switch d {
	case "L":
		gomos[0] -= 1
	case "R":
		gomos[0] += 1
	case "U":
		gomos[1] -= 1
	case "D":
		gomos[1] += 1
	}
	for i := 0; i < n*2; i += 2 {
		fmt.Printf("%d %d\n", gomos[i], gomos[i+1])
	}
}

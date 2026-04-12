package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func getMen(vet []int) []int {
	men := make([]int, 0, len(vet))
	for i := 0; i < len(vet); i++ {
		if vet[i] > 0 {
			men = append(men, vet[i])
		}
	}
	return men
}

func getCalmWomen(vet []int) []int {
	woman := make([]int, 0, len(vet))
	for i := 0; i < len(vet); i++ {
		if vet[i] < 0 && vet[i] > -10 {
			woman = append(woman, vet[i])
		}
	}
	return woman
}

func sortVet(vet []int) []int {
	slices.Sort(vet)
	return vet
}

func sortStress(vet []int) []int {
	counts := make(map[int]int)
	final := make([]int, 0, len(vet))

	for _, n := range vet {
		v := n
		if n < 0 {
			v = -n
			counts[v]++
		}
		final = append(final, v)
	}

	slices.Sort(final)

	for i, v := range final {
		if counts[v] > 0 {
			final[i] = -v
			counts[v]--
		}
	}

	return final
}

func reverse(vet []int) []int {
	res := slices.Clone(vet)
	slices.Reverse(res)
	return res
}

func unique(vet []int) []int {
	m := make(map[int]bool)
	res := make([]int, 0)
	for _, v := range vet {
		if !m[v] {
			m[v] = true
			res = append(res, v)
		}
	}
	return res
}

func repeated(vet []int) []int {
	vistos := make(map[int]bool)
	res := []int{}

	for _, v := range vet {
		if vistos[v] {
			res = append(res, v)
		} else {
			vistos[v] = true
		}
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if !scanner.Scan() {
			break
		}
		fmt.Print("$")
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "get_men":
			printVec(getMen(str2vet(args[1])))
		case "get_calm_women":
			printVec(getCalmWomen(str2vet(args[1])))
		case "sort":
			printVec(sortVet(str2vet(args[1])))
		case "sort_stress":
			printVec(sortStress(str2vet(args[1])))
		case "reverse":
			array := str2vet(args[1])
			other := reverse(array)
			printVec(array)
			printVec(other)
		case "unique":
			printVec(unique(str2vet(args[1])))
		case "repeated":
			printVec(repeated(str2vet(args[1])))
		case "end":
			return
		}
	}
}

func printVec(vet []int) {
	fmt.Print("[")
	for i, val := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(val)
	}
	fmt.Println("]")
}

func str2vet(s string) []int {
	if s == "[]" {
		return nil
	}
	s = s[1 : len(s)-1]
	parts := strings.Split(s, ",")
	var vet []int
	for _, part := range parts {
		n, _ := strconv.Atoi(part)
		vet = append(vet, n)
	}
	return vet
}

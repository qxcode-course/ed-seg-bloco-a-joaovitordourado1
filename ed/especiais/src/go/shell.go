package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"slices"
	"cmp"
)

type Pair struct {
	One int
	Two int
}

func occurr(vet []int) []Pair {
	res := []Pair{}
	repetidas := make(map[int]int)

	for i := 1; i <= len(vet); i++ {
		if vet[i] < 0{
			vet[i]*= -1
		}
		if(repetidas[vet[i]]==0){
			repetidas[vet[i]]=1
		}else{
			repetidas[vet[i]]++
		}
	}
	for i ,ele := range repetidas{
		res = append(res,Pair{i , ele})
	}
	slices.SortFunc(res, func(a,b Pair) int {return cmp.Compare(a.One, b.One)})
	return res
}

func teams(vet []int) []Pair {
	if len(vet) == 0 {
		return nil
	}
	res := []Pair{}
	count := 1
	for i := 1; i <= len(vet); i++ {
		if i == len(vet) || vet[i] != vet[i-1] {
			res = append(res, Pair{vet[i-1], count})
			count = 1
		} else {
			count++
		}
	}
	return res
}

func mnext(vet []int) []int {
	_ = vet
	return nil
}

func alone(vet []int) []int {
	_ = vet
	return nil
}

func couple(vet []int) int {
	_ = vet
	return 0
}

func hasSubseq(vet []int, seq []int, pos int) bool {
	_ = vet
	_ = seq
	_ = pos
	return false
}

func subseq(vet []int, seq []int) int {
	_ = vet
	_ = seq
	return -1
}

func erase(vet []int, posList []int) []int {
	_ = vet
	_ = posList
	return nil
}

func clear(vet []int, value int) []int {
	_ = vet
	_ = value
	return nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "occurr":
			printSlice(occurr(str2vet(args[1])))
		case "teams":
			printSlice(teams(str2vet(args[1])))
		case "mnext":
			printSlice(mnext(str2vet(args[1])))
		case "alone":
			printSlice(alone(str2vet(args[1])))
		case "erase":
			printSlice(erase(str2vet(args[1]), str2vet(args[2])))
		case "clear":
			val, _ := strconv.Atoi(args[2])
			printSlice(clear(str2vet(args[1]), val))
		case "subseq":
			fmt.Println(subseq(str2vet(args[1]), str2vet(args[2])))
		case "couple":
			fmt.Println(couple(str2vet(args[1])))
		case "end":
			return
		default:
			fmt.Println("Invalid command")
		}
	}
}

// Funções auxiliares

func str2vet(str string) []int {
	if str == "[]" {
		return nil
	}
	str = str[1 : len(str)-1]
	parts := strings.Split(str, ",")
	var vet []int
	for _, part := range parts {
		num, _ := strconv.Atoi(strings.TrimSpace(part))
		vet = append(vet, num)
	}
	return vet
}

func printSlice[T any](vet []T) {
	fmt.Print("[")
	for i, x := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(x)
	}
	fmt.Println("]")
}

func (p Pair) String() string {
	return fmt.Sprintf("(%v, %v)", p.One, p.Two)
}

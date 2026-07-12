package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type multiSet struct {
	data     []int
	size     int
	capacity int
}

func NewmultiSet(v int) *multiSet {
	return &multiSet{data: make([]int, v),
		size:     0,
		capacity: v,
	}
}
func (s *multiSet) expand() {
	if s.capacity == 0 {
		s.capacity = 1
	} else {
		s.capacity *= 2
	}
	newData := make([]int, s.capacity)
	for i := 0; i < s.size; i++ {
		newData[i] = s.data[i]
	}
	s.data = newData
}
func (s *multiSet) Contains(b int) bool {
	for i := 0; i < s.size; i++ {
		if s.data[i] == b {
			return true
		}
	}
	return false

}
func (s *multiSet) Count(b int) int {
	cont := 0
	for i := 0; i < s.size; i++ {
		if s.data[i] == b {
			cont++
		}
	}
	return cont

}
func (s *multiSet) Unique() int {
	u := make(map[int]bool)
	for i := 0; i < s.size; i++ {
		u[s.data[i]] = true
	}
	return len(u)
}

func (s *multiSet) Clear(){
	s.size =0
}

func (s *multiSet) search(b int) (bool, int) {
	h := 0
	for i := 0; i < s.size; i++ {
		if b == s.data[i] {
			h = i
		}
	}
	if h == 0 {
		return false, -1
	}
	return true, h
}
func (s *multiSet) Insert(v int) {
	if s.size == s.capacity {
        s.expand()
    }
	aux := 0
	aux1 := -1

	for i := 0; i < s.size; i++ {
		if s.data[i] >= v {
			aux = s.data[i]
			aux1 = i
			s.data[i] = v
			break
		}
	}
	if aux1 != -1 {
		for i := aux1 + 1; i <=s.size; i++ {
			temp := s.data[i]
			s.data[i] = aux
			aux = temp
		}
	} else {
		s.data[s.size]=v
	}
	s.size++
}
func (s *multiSet) insert(v int, index int) error {
	for i := s.size; i > index; i-- {
		s.data[i] = s.data[i-1]
	}
	s.data[index] = v
	s.size++
	return nil
}
func (s *multiSet) Erase(v int) error {
	for i := 0; i < s.size; i++ {
		if s.data[i] == v {
			for j := i; j < s.size-1; j++ {
                s.data[j] = s.data[j+1]
            }
			s.size--
			return nil
		}
	}
	return errors.New("value not found")

}
func (s *multiSet) erase(v int, index int) error {
	if index < 0 || index >= s.size {
		return errors.New("índice fora dos limites")
	}
	for i := index; i < s.size-1; i++ {
		s.data[i] = s.data[i+1]
	}
	s.size--
	return nil
}

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	result := fmt.Sprintf("%d", slice[0])
	for _, value := range slice[1:] {
		result += sep + fmt.Sprintf("%d", value)
	}
	return result
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)
	ms := NewmultiSet(0)

	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		args := strings.Fields(line)
		fmt.Println(line)
		if len(args) == 0 {
			continue
		}
		cmd = args[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(args[1])
			ms = NewmultiSet(value)
		case "insert":
			for _, part := range args[1:] {
				value, _ := strconv.Atoi(part)
				ms.Insert(value)
			}
		case "show":
			fmt.Printf("[%s]\n", Join(ms.data[:ms.size], ", "))
		case "erase":
			value, _ := strconv.Atoi(args[1])
			err := ms.Erase(value)
            if err != nil {
                fmt.Println(err)
            }
		case "contains":
			value, _ := strconv.Atoi(args[1])
			fmt.Println(ms.Contains(value))
		case "count":
			value, _ := strconv.Atoi(args[1])
            fmt.Println(ms.Count(value))
		case "unique":
			fmt.Println(ms.Unique())
		case "clear":
			ms.Clear()
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
// tive que dar uma alterada na main, usei ia para isso pois nao tinha conhecimento da lib de conversao
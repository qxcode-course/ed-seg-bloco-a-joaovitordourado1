package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
type multiSet struct {
	data []int
	size int
	capacity int

}
func NewmultiSet(v int) *multiSet {
	return &multiSet{data: make([]int,v),
	size: 0,
	capacity: v,
	}
}
func (s *multiSet) expand() {
	if s.capacity ==0{
		s.capacity = 1
	}else{
		s.capacity *=2
	}
	newData := make([]int,s.capacity )
	for i := 0; i < s.size; i++ {
		newData[i] = s.data[i]
	}
	s.data = newData
}
func (s *multiSet) search(){
	
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
	// ms := NewMultiSet(0)

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
			// value, _ := strconv.Atoi(args[1])
			// ms = NewMultiSet(value)
		case "insert":
			// for _, part := range args[1:] {
			// 	value, _ := strconv.Atoi(part)
			// }
		case "show":
		case "erase":
			// value, _ := strconv.Atoi(args[1])
		case "contains":
			// value, _ := strconv.Atoi(args[1])
		case "count":
			// value, _ := strconv.Atoi(args[1])
		case "unique":
		case "clear":
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}

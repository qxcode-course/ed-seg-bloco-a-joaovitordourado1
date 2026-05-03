package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	
)

type Set struct {
	data []int
	size int
	capacity int

}

func NewSet(v int) *Set {
	return &Set{data: make([]int,v),
	size: 0,
	capacity: v,
	}
}
func (s *Set) insert(n int, index int) error {
	if s.size == s.capacity {
		s.Reserve(s.capacity*2)
	}
	for i := s.size; i > index; i-- {
		s.data[i] = s.data[i-1]
	}
	s.data[index] = n
	s.size++
	
	return nil
}
func (s *Set) Insert(v int) {
	if !s.Contains(v){
		low := 0 
	high:=s.size-1
	for low<=high{
		mid := low + (high - low)/2 
		if s.data[mid] < v{
			low = mid +1
		}else{
			high = mid -1
		}
		
	}
	s.insert(v,low)
	}
	
}

func (s *Set) Contains(v int) bool{
	k:=s.binarySearch(v)
	if k ==-1{
		return false
	}
	return true
}

func (s *Set) Reserve(newCapacity int) {
	newData := make([]int, newCapacity)
	for i := 0; i < s.size; i++ {
		newData[i] = s.data[i]
	}
	s.data = newData
	s.capacity = newCapacity
}
func (s *Set) erase(id int) error{
	for i := id; i < s.size-1; i++ {
		s.data[i] = s.data[i+1]
	}
	s.size--
	return nil
}
func (s *Set) Erase(num int ) bool{
	if!s.Contains(num){
		return false
	}
	indva := s.binarySearch(num)

	s.erase(indva)
	return true 
}
func (s *Set) binarySearch(v int) int{
	low := 0 
	high:=s.size-1
	for low<=high{
		mid := low + (high - low)/2 
		if s.data[mid]== v {
			return mid
		}
		if s.data[mid] < v{
			low = mid +1
		}else{
			high = mid -1
		}
	}
	return -1
}
func (s *Set) String() string {
	jap := make([]string, s.size)
	for i := 0; i < s.size; i++ {
		jap[i] = strconv.Itoa(s.data[i])
	}
	return "[" + strings.Join(jap, ", ") + "]"
}
func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)
	v := NewSet(0)
	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		fmt.Println(line)
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		cmd = parts[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(parts[1])
			v = NewSet(value)
		case "insert":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				v.Insert(value)
			}
		case "show":
            fmt.Println(v.String())

		case "erase":
			value, _ := strconv.Atoi(parts[1])
			if !v.Erase(value){
				fmt.Println("value not found")
			}
		case "contains":
			value, _ := strconv.Atoi(parts[1])
			fmt.Println(v.Contains(value))
		case "clear":
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Vector struct {
	data     []int
	size     int
	capacity int
}

func NewVector(capacity int) *Vector {
	return &Vector{
		data:     make([]int, capacity), // nunca use len(data) ou cap(data) ou qq método do go de manipulação de array
		size:     0,
		capacity: capacity,
	}
}
func (v *Vector) Reserve(newCapacity int) {
	if v.capacity >= newCapacity {
		return
	}
	newData := make([]int, newCapacity)
	for i := 0; i < v.size; i++ {
		newData[i] = v.data[i]
	}
	v.data = newData
	v.capacity = newCapacity
}
func (v *Vector) PushBack(value int) {
	if v.size == v.capacity {
		a := v.capacity * 2
		if a == 0 {
			a = 1
		}
		v.Reserve(a)
	}
	v.data[v.size] = value
	v.size++
}
func (v *Vector) Size() int {
	return v.size
}
func (v *Vector) PopBack() (int, error) {
	if v.size == 0 {
		return 0, errors.New("vector is empty")
	}
	num := v.size - 1
	removido := v.data[num]

	v.size--

	return removido, nil
}
func (v *Vector) Capacity() int {
	return v.capacity
}

func (v *Vector) Status() string {
	return fmt.Sprint("size:", v.size, " capacity:", v.capacity)
}

func (v *Vector) String() string {
	if v.size == 0 {
		return "[]"
	}
	return "[" + Join(v.data[:v.size], ", ") + "]"
}
func (v *Vector) Get(n int) int {
	return v.data[n]
}
func (v *Vector) At(n int) (int, error) {
	if n >= v.size {
		return -1, errors.New("index out of range")
	}
	return v.Get(n), nil
}
func (v *Vector) Set(n int, va int) error {
	if n >= v.size {
		return errors.New("index out of range")
	}
	v.data[n] = va
	return nil
}
func (v *Vector) Erase(id int) error {
	if id >= v.size {
		return errors.New("index out of range")
	}
	for i := id; i < v.size-1; i++ {
		v.data[i] = v.data[i+1]
	}
	v.size--
	return nil
}
func (v *Vector) Slice(s int, e int) *Vector{
	k:=(s%v.size+v.size)%v.size
	p:=(e%v.size+v.size)%v.size
	f:= v.data[k:p]
	if k>p{
		p = v.size
	}
	return &Vector{
		data : f,
		size : len(f),
		capacity : cap(f),
	}


}
func (v *Vector) Insert(n int, va int) error {
	if v.size == v.capacity {
		v.Reserve(v.capacity*2)
	}
	for i := v.size; i > n; i-- {
		v.data[i] = v.data[i-1]
	}
	v.size++
	v.data[n] = va
	return nil
}
func (v *Vector) Clear(){
	v.size = 0
}
func (v *Vector) IndexOf(k int) int{
	for i := 0; i < v.size; i++ {
		if k == v.data[i]{
			return i 
		}
	}
	return -1
}
func (v *Vector) Contains(p int) bool{
	for i := 0; i < v.size; i++ {
		if v.data[i] == p{
			return true
		}
	}
	return false
}
func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	var result strings.Builder
	fmt.Fprintf(&result, "%d", slice[0])
	for _, value := range slice[1:] {
		fmt.Fprintf(&result, "%s%d", sep, value)
	}
	return result.String()
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	v := NewVector(0)
	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
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
			v = NewVector(value)
		case "push":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				v.PushBack(value)
			}
		case "show":
			fmt.Println(v.String())
		case "status":
			fmt.Println(v.Status())
		case "pop":
			_, err := v.PopBack()
			if err != nil {
				fmt.Println(err)
			}
		case "insert":
			index, _ := strconv.Atoi(parts[1])
			value, _ := strconv.Atoi(parts[2])
			err := v.Insert(index, value)
			if err != nil {
				fmt.Println(err)
			}
		case "erase":
			index, _ := strconv.Atoi(parts[1])
			err := v.Erase(index)
			if err != nil {
				fmt.Println(err)
			}
		case "indexOf":
			value, _ := strconv.Atoi(parts[1])
			index := v.IndexOf(value)
			fmt.Println(index)
		case "contains":
			value, _ := strconv.Atoi(parts[1])
			if v.Contains(value) {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}
		case "clear":
			v.Clear()
		case "capacity":
			fmt.Println(v.Capacity())
		case "get":
			index, _ := strconv.Atoi(parts[1])
			value, err := v.At(index)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(value)
			}
		case "set":
			index, _ := strconv.Atoi(parts[1])
			value, _ := strconv.Atoi(parts[2])
			err := v.Set(index, value)
			if err != nil {
				fmt.Println(err)
			}
			
		case "reserve":
			newCapacity, _ := strconv.Atoi(parts[1])
			v.Reserve(newCapacity)
		case "slice":
			start, _ := strconv.Atoi(parts[1])
			end, _ := strconv.Atoi(parts[2])
			slice := v.Slice(start, end)
			fmt.Println(slice)
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}

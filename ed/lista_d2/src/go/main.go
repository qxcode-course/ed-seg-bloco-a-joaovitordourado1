package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
type Node struct {
	Value int
	next  *Node
	prev  *Node
	root *Node
}
func(n *Node)Next() *Node{
	return n.next
}
func(n *Node)Prev() *Node{
	return n.prev
}
type LList struct{
	root *Node
	size int 
}
func (l *LList) Size() int{
	return l.size
}
func (l *LList) Clear(){
	l.root.prev = l.root
	l.root.next = l.root
}
func (n *LList) PushFront(k int) {
	newno := &Node{
		Value: k,
		next:  n.root.next,
		prev:  n.root,
	}
	n.root.next = newno
	newno.next.prev = newno
}
func (n *LList) PushBack(k int) {
	newno := &Node{
		Value: k,
		next:  n.root,
		prev:  n.root.prev,
	}
	n.root.prev = newno
	newno.prev.next = newno
}
func (n *LList) PopFront() {
	if n.root.next != n.root {
		n.root.next = n.root.next.next
		n.root.next.prev = n.root
	}

}
func (n *LList) PopBack() {
	if n.root.prev != n.root {
		n.root.prev = n.root.prev.prev
		n.root.prev.next = n.root
	}

}
func (n *LList) Front() *Node{
	return n.root.Next()
}
func(n *LList) Back() *Node{
	return n.root.Prev()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
// 	ll := NewLList()

	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Println(line)
		args := strings.Fields(line)

		if len(args) == 0 {
			continue
		}

		cmd := args[0]

		switch cmd {
		case "show":
			// fmt.Println(ll.String())
		case "size":
			// fmt.Println(ll.Size())
		case "push_back":
			// for _, v := range args[1:] {
			// 	num, _ := strconv.Atoi(v)
			// 	ll.PushBack(num)
			// }
		case "push_front":
			// for _, v := range args[1:] {
			// 	num, _ := strconv.Atoi(v)
			// 	ll.PushFront(num)
			// }
		case "pop_back":
			// ll.PopBack()
		case "pop_front":
			// ll.PopFront()
		case "clear":
			// ll.Clear()
		case "walk":
			// fmt.Print("[ ")
			// for node := ll.Front(); node != nil; node = node.Next() {
			// 	fmt.Printf("%v ", node.Value)
			// }
			// fmt.Print("]\n[ ")
			// for node := ll.Back(); node != nil; node = node.Prev() {
			// 	fmt.Printf("%v ", node.Value)
			// }
			// fmt.Println("]")
		case "replace":
			// oldvalue, _ := strconv.Atoi(args[1])
			// newvalue, _ := strconv.Atoi(args[2])
			// node := ll.Search(oldvalue)
			// if node != nil {
			// 	node.Value = newvalue
			// } else {
			// 	fmt.Println("fail: not found")
			// }
		case "insert":
			// oldvalue, _ := strconv.Atoi(args[1])
			// newvalue, _ := strconv.Atoi(args[2])
			// node := ll.Search(oldvalue)
			// if node != nil {
			// 	ll.Insert(node, newvalue)
			// } else {
			// 	fmt.Println("fail: not found")
			// }
		case "remove":
			// oldvalue, _ := strconv.Atoi(args[1])
			// node := ll.Search(oldvalue)
			// if node != nil {
			// 	ll.Remove(node)
			// } else {
			// 	fmt.Println("fail: not found")
			// }
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}

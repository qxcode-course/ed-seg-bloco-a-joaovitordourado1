package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Value int
	next  *Node
	prev  *Node
}
type llist struct {
	root *Node
}

func (n *llist) Size() int {
	cont := 0
	massa := n.root.next
	for massa != n.root {
		cont++
		massa = massa.next
	}
	return cont
}
func (n *llist) Clear() {
	n.root.prev = n.root
	n.root.next = n.root
}
func (n *llist) PushFront(k int) {
	newno := &Node{
		Value: k,
		next:  n.root.next,
		prev:  n.root,
	}
	n.root.next = newno
	newno.next.prev = newno
}
func (n *llist) PushBack(k int) {
	newno := &Node{
		Value: k,
		next:  n.root,
		prev:  n.root.prev,
	}
	n.root.prev = newno
	newno.prev.next = newno
}
func (n *llist) PopFront() {
	if n.root.next != n.root {
		n.root.next = n.root.next.next
		n.root.next.prev = n.root
	}

}
func (n *llist) PopBack() {
	if n.root.prev != n.root {
		n.root.prev = n.root.prev.prev
		n.root.prev.next = n.root
	}

}
func (n *llist) String() string {
	saida := "["
	massa := n.root.next

	for massa != n.root {
		saida += fmt.Sprintf("%v", massa.Value)

		if massa.next != n.root {
			saida += ", "
		}
		massa = massa.next
	}
	saida += "]"
	return saida
}
func NewLList() *llist {
	root := &Node{}

	root.next = root
	root.prev = root
	return &llist{
		root: root,
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ll := NewLList()

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
			fmt.Println(ll.String())
		case "size":
			fmt.Println(ll.Size())
		case "push_back":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushBack(num)
			}
		case "push_front":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushFront(num)
			}
		case "pop_back":
			ll.PopBack()
		case "pop_front":
			ll.PopFront()
		case "clear":
			ll.Clear()
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}

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
	Left  *Node
	Right *Node
}

func BstInsert(values []int) *Node {
	var k *Node
	for _,value :=range values{
	if k == nil{
		k=&Node{Value:value}
		continue
	}
	ac := k
	for{
		if value < ac.Value{
			if ac.Left == nil{
				ac.Left = &Node{Value : value}
				break

			}
			ac = ac.Left
			}else if value > ac.Value{
				if ac.Right == nil{
					ac.Right = &Node{Value : value}
					break
				}
				ac = ac.Right
			}else{
				break
			}
		}



	}
	return k
	
}
// Dica: crie um vetor compartilhado e vá preenchendo conforme anda na recursão
// Depois use o strings.Join para gerar o serial
func Serialize(root *Node) string {
	if root== nil{
		return "#"
	}
	esq:=Serialize(root.Left)
	rig:=Serialize(root.Right)

	return fmt.Sprintf("%d %s %s", root.Value,esq,rig)
}

// -----------------------------------------------------------------------------------
func BShow(node *Node, history string) {
	if node != nil && (node.Left != nil || node.Right != nil) {
		BShow(node.Left, history+"l")
	}
	for i := 0; i < len(history)-1; i++ {
		if history[i] != history[i+1] {
			fmt.Print("│   ")
		} else {
			fmt.Print("    ")
		}
	}
	if history != "" {
		if history[len(history)-1] == 'l' {
			fmt.Print("╭───")
		} else {
			fmt.Print("╰───")
		}
	}
	if node == nil {
		fmt.Println("#")
		return
	}
	fmt.Println(node.Value)
	if node.Left != nil || node.Right != nil {
		BShow(node.Right, history+"r")
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Split(scanner.Text(), " ")
	values := make([]int, 0, len(parts))
	for _, elem := range parts {
		v, err := strconv.Atoi(elem)
		if err == nil {
			values = append(values, v)
		}
	}
	root := BstInsert(values)
	BShow(root, "") // Chama a função de impressão formatada
	fmt.Println(Serialize((root)))
}

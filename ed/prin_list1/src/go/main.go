package main

import (
	"fmt"
)


// mostra a lista com o elemento sword destacado
func ToStr(l *DList[int], sword *DNode[int]) string {
	k:= "[ "
	for i:= l.root.next; i!=l.root;i=i.next{
		if i!= sword{
		k += fmt.Sprintf("%v ", i.Value)
		}
		if i == sword{
			k += fmt.Sprintf("%v", i.Value)
			k += "> "
		}
	}
	k +="]"
	return k
}

// move para frente na lista circular
func Next(l *DList[int], it *DNode[int]) *DNode[int] {
	prox := it.next

	if prox == l.root{
		prox = prox.next
	}
	return prox
	
}

func main() {
	var qtd, chosen int
	fmt.Scan(&qtd, &chosen)
	l := NewDList[int]()
	for i := 1; i <= qtd; i++ {
		l.PushBack(i)
	}
	sword := l.Front()
	for range chosen - 1 {
		sword = Next(l, sword)
	}
	for range qtd - 1 {
		fmt.Println(ToStr(l, sword))
		l.Erase(Next(l, sword))
		sword = Next(l, sword)
	}
	fmt.Println(ToStr(l, sword))
}

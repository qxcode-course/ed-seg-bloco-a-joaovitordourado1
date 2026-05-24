package main

import (
	"container/list"
	"fmt"
)

// mostra a lista com o elemento sword destacado
func ToStr(l *list.List, sword *list.Element) string {
	k:="[ "
	for i := l.Front(); i != nil; i = i.Next() {
        if i != sword {
            k += fmt.Sprintf("%v ", i.Value)
        } else {
            val := i.Value.(int)
            if val > 0 {
                k += fmt.Sprintf("%v> ", val)
            } else {
                k += fmt.Sprintf("<%v ", val) 
            }
        }
    }
	k += "]"
	return k
}

// move para frente na lista circular
func Next(l *list.List, it *list.Element) *list.Element {
	prox := it.Next()
	if prox == nil{
		prox = l.Front()
	}
	return prox
}

// move para tras na lista circular
func Prev(l *list.List, it *list.Element) *list.Element {
	prev := it.Prev()
	if prev == nil{
		prev = l.Back()
	}
	return prev
}

func main() {
	var qtd, chosen, fase int
	fmt.Scan(&qtd, &chosen, &fase)
	l := list.New()
	for i := 1; i <= qtd; i++ {
		l.PushBack(i * fase)
		fase = -fase
	}
	sword := l.Front()
	for range chosen - 1 {
		sword = Next(l, sword)
	}
	for range qtd - 1 {
		fmt.Println(ToStr(l, sword))
		if sword.Value.(int) > 0 {
			l.Remove(Next(l, sword))
			sword = Next(l, sword)
		} else {
			l.Remove(Prev(l, sword))
			sword = Prev(l, sword)
		}
	}
	fmt.Println(ToStr(l, sword))
}

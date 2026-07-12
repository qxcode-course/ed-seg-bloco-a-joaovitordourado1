package main
import "fmt"
type deque struct{
    elements[]rune 
}


func(d *deque )pushBack(i rune) {
    d.elements = append(d.elements,i)
}
func(d *deque)popBack()(rune){
    i := len(d.elements)-1
    x := d.elements[i]
    d.elements = d.elements[:i]
    return x
}
func tabal(s []rune) bool{
    d := deque {}
    for _, c := range s{
        if c =='(' || c == '['{
            d.pushBack(c)
        }else if c == ')' || c == ']'{
            if len(d.elements)== 0{
                return false 
            }
            t:=d.popBack()
            if (c == ')' && t!= '(')||(c == ']' && t!= '['){
                return false
            }
        }
    }

    return len(d.elements) == 0
}
func main() {
    in := ""
    fmt.Scan(&in)
    k := []rune(in)
    if tabal(k){
        fmt.Println("balanceado")
    }else{
        fmt.Println("nao balanceado")
    }
    
}

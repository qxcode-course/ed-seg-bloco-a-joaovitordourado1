package main
import "fmt"
func main() {
    cont :=0
    n:=0
    fmt.Scan(&n)
    div :=2
    enes(n,cont,div)
}
func enes(n int, cont int,atual int ){
    if eh_primo(atual,2){
        cont++
        if n == cont {
        fmt.Println(atual)
        return
    }
    }
    
    enes(n,cont,atual+1)
}
func eh_primo(x int, div int) bool {
	if x < 2 {
		return false
	}
	if div == x {
		return true
	}
	if x%div == 0 {
		return false
	}
	if div > x {
		return true
	}
	return eh_primo(x, div+1)
}

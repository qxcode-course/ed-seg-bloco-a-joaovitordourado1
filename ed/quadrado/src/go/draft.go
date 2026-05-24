package main

import "fmt"

func quadrado(n int) string {

	k := ""
	if n == 1 {
        k+="1^2 = 1"
        return k +"\n"
    }
    p:=fmt.Sprintf("%v^2 = %v^2 + 2*%v + 1 = ?\n",n,n-1,n-1)
    prox:=quadrado(n-1)
    r:=fmt.Sprintf("%v^2 = %v^2 + 2*%v + 1 = %v\n",n,n-1,n-1,n*n)

    return p + prox +r
}
func main() {
	n := 0
	fmt.Scan(&n)
    resultado := quadrado(n)
    fmt.Print(resultado)
}

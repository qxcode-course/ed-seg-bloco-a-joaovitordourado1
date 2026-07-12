package main
import "fmt"
func main() {
    p:=0
    v := []int{}
    fmt.Scan(&p)
    w:=xx(v,p,2)
    fmt.Print("[")
    for i:=0; i < len(w);i++{
        if i!=len(w)-1{
        fmt.Printf("%v, ",w[i])
        }else {
        fmt.Printf("%v",w[i])
    }}
     fmt.Println("]")
}
func xx(k []int ,n int, atual int)[]int{
    if len(k) == n{
        return k 
    }
    if  epri(atual){
        k = append(k,atual)
    }
    return xx(k,n,atual+1)
}
func epri(b int) bool{
    if b ==1{
        return false
    }
    for i:= 2; i< b; i++{
        if b%i ==0{
            return false
        }
    } 
    return true
}
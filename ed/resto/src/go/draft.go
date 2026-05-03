package main
import "fmt"
func main() {
    n:=0 
    fmt.Scan(&n)
    Sigma(n)
}
func Sigma(n int){
    if n ==1{
        fmt.Println("0 1")
        return
    }
    Sigma((n)/2)
    if n%2 ==1{
        fmt.Println((n)/2 ,"1")
    } else{
        fmt.Println( (n)/2 ,"0" )
    }
}

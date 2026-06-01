package main
import "fmt"

                                            
func main() {
    n:=0
    k:=0
    fmt.Scan(&n)
    vec:= make([]int,n*2)                                                     
    for i := 0; i < n*2; i++ {
        fmt.Scan(&vec[i])
    }
    for i := 0; i < n*2; i++ {
        if vec[i] < vec[i+1]{
            continue 
        }
        if vec[i] > vec[i+1]{
            k = vec[i] - vec[i+1]
        }

    }
    fmt.Println(k)
}
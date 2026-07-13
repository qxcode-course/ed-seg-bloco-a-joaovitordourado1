package main
import "fmt"

                                            
func main() {
    n:=0
    fmt.Scan(&n)
    vec:= make([]int,n)     
    gas:= make([]int,n)                                             
    in:=0
    ta:=0
    sg:=0                      
    for i := 0; i < n; i++ {
        fmt.Scan(&gas[i],&vec[i])
    }
    for i := 0; i < n; i++{
        sobra := gas[i] - vec[i]
        ta+= sobra
        sg +=sobra
        if ta<0{
            ta =0
            in = i+1
        }
    }
    if sg>=0{
        fmt.Println(in)    }else{
            fmt.Println("-1")
        }}

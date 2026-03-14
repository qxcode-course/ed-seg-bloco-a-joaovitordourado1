package main 
import ("fmt"
)
func Ladrao(h int,p int, f int , d int ) string{
    for{
        if f == p{
            return "N"
        }
        if f == h{
            return "S"
        }
        f = (f+d+16)%16
    }
} 
func main(){
    var h,p,f,d int
    fmt.Scan(&h,&p,&f,&d)
    valor := Ladrao(h,p,f,d)
    fmt.Println(valor)
}

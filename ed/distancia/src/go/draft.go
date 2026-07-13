package main
import "fmt"

func aux(a[]byte,id,L int,char byte)bool{
    for i := 1; i <= L; i++ {
        if id-i>=0 && a[id-i]== char{
            return false
        }
    }
    for i := 1; i <=L; i++ {
        if id+i< len(a)&& a[id+i]== char{
            return false
        }
    }
    return true
}
func back(a []byte,L,i int)bool{
    if i == len(a){
        return true
    }
    if a[i] != '.'{
        return back(a,L,i+1)
    }
    for d := 0; d <= L; d++{
        cd:= byte('0'+d)
        if aux(a,i,L,cd){
            a[i]=cd

        if back(a,L,i){
            return true
        }
             a[i]= '.'
        }
    }
    return false
}
func main() {
    s := ""
    L:=0

    fmt.Scan(&s)
    fmt.Scan(&L)


    a:=[]byte(s)
    back(a,L,0)
    fmt.Println(string(a))
}
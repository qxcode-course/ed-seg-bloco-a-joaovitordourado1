package main

import ("fmt" 
        "math"
        )
func main(){
    var n1, n2, n3, n4, n5  float64
    fmt.Scan(&n1, &n2, &n3)
    n4 = (n1+n2+n3)/2
    n5 =math.Sqrt(n4*(n4-n1)*(n4-n2)*(n4-n3))
    fmt.Printf("%.2f\n",n5)
}
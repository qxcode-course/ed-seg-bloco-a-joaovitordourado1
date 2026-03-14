package main 
import("fmt"
    )
    
    func calcIdade(nome string, idade int ) string {
        if idade<12{
            return nome +" eh crianca\n"
        }  else if idade<18 {
            return nome +" eh jovem\n"
        }  else if idade<65 {
            return nome + " eh adulto\n"
        }  else if idade <1000{
            return nome + " eh idoso\n"
        }  else{
            return nome + " eh mumia\n"
        }
    }
    func main(){
        var nome string
        var idade int
        var frase string
        fmt.Scan(&nome,&idade)
        frase = calcIdade(nome,idade)
        fmt.Printf(frase)
    }
    
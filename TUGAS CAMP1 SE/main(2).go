package main 

import (
    "fmt"
)

func main () {
    fmt.Print("masukan panjang alas segitiga : ")
    var alas float64
    fmt.Scanln(&alas)
    
    fmt.Print("masukan tinggi segitiga : ")
    var tinggi float64
    fmt.Scanln(&tinggi)
    
    luas := 0.5 * alas * tinggi 
    
    fmt.Print("luas segitiga adalah %.2f", luas)

}

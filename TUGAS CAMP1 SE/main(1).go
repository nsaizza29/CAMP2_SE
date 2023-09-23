package main 

import (
    "fmt"
)

func main () {
    fmt.Print("masukan panjang sisi : ")
    var sisi float64
    fmt.Scanln(&sisi)
    luas := sisi * sisi
    fmt.Printf("luas persegi dengan sisi %.2f adalah %.2f", sisi,luas)

}

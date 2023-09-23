package main 

import (
    "fmt"
    "math"
)

func main () {
    fmt.Print("masukan jari-jari lingkaran : ")
    var jarijari float64
    fmt.Scanln(&jarijari)
    
    luas := math.Pi * jarijari * jarijari 
    
    fmt.Print("luas lingkaran adalah %.2f", luas)

}

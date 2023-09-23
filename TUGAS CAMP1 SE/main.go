package main 

import (
    "fmt"
)

func main () {
    fmt.Print("masuksan periode getaran (detik): ")
    var periode float64
    fmt.Scanln(&periode)
    
    frekuensi := 1.0 / periode
    
    fmt.Printf("frequensi getaran adalah %.2f hz\n", frekuensi)

}

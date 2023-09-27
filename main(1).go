/******************************************************************************
Welcome to GDB Online.
GDB online is an online compiler and debugger tool for C, C++, Python, Java, PHP, Ruby, Perl,
C#, OCaml, VB, Swift, Pascal, Fortran, Haskell, Objective-C, Assembly, HTML, CSS, JS, SQLite, Prolog.
Code, Compile, Run and Debug online from anywhere in world.

*******************************************************************************/
package main
import "fmt"

func main() {
    data := []int{16, 11, 7, 9, 32, 1, 26, 5, 31, 11, 3}
    
    bubbleSort(data)
    
    fmt.Println("data terurut dari terbesar ke terkecil :", data)
}

func bubbleSort(arr []int) {
    n := len(arr)
    for i := 0; i < n-1; i++ {
        for j := 0; j < n-i-1; j++ {
            //jika elemen saat ini lebih kecil dari elemen berikutnnya, tukar
            if arr[j] < arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }    
        }    
    }
}

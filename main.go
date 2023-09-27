/******************************************************************************
Welcome to GDB Online.
GDB online is an online compiler and debugger tool for C, C++, Python, Java, PHP, Ruby, Perl,
C#, OCaml, VB, Swift, Pascal, Fortran, Haskell, Objective-C, Assembly, HTML, CSS, JS, SQLite, Prolog.
Code, Compile, Run and Debug online from anywhere in world.

*******************************************************************************/
package main
import "fmt"

func main() {
    data := []int{1, 2, 3, 4, 5}
    fmt.Println("data awal : ", data)
    data = append(data, 6, 7, 8)
    
    fmt.Println("setelah penambahan:", data)
}
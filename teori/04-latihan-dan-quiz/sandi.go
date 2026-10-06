package main

import "fmt"

func main() {
    var n, i, bilangan, jumlah, digit1, digit4 int
    fmt.Scan(&n)

    for i = 1; i <= n; i++ {
        fmt.Scan(&bilangan)
        digit1 = bilangan / 1000
        digit4 = bilangan % 10
        jumlah += digit1 + digit4
    }  
    fmt.Println(jumlah)
}
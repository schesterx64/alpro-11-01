package main

import "fmt"

func main() {
	var x float64

	fmt.Println("Masukkan jarak dalam mil")
	fmt.Scan(&x)

	kilometer := x * 1.6

	fmt.Println("Konversi ke kilometer", kilometer)

}
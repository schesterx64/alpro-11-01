package main 

import "fmt"

func main() {
	fmt.Println("Masukkan suhu dalam Celcius:")
	var C float64
	fmt.Scan(&C)

	fmt.Println("Konversi ke Reamur:", (4.0/5.0) * C)
}
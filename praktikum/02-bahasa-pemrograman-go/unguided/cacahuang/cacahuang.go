package main

import "fmt"

func main() {
	var uang int
	fmt.Scan(&uang)

	//Uang dalam lembar 10.000
	sepuluhRibu := uang / 10000
	sisa := uang % 10000

	//Uang dalam lembar 5.000
	limaRibu := sisa / 5000
	sisa = sisa % 5000

	//Uang dalam lembar 1000
	seRibu := sisa / 1000

	fmt.Printf("Sepuluh Ribu = %d, Lima Ribu = %d, Seribu = %d\n", sepuluhRibu, limaRibu, seRibu)
}

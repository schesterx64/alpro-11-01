package main

import "fmt"

func main() {
	var uang int
	fmt.Scan(&uang)

	sepuluhRibu := uang / 10000
	sisa := uang % 10000

	limaRibu := sisa / 5000
	sisa = sisa % 5000

	seRibu := sisa / 1000

	fmt.Printf("Sepuluh Ribu = %d, Lima Ribu = %d, Seribu = %d\n", sepuluhRibu, limaRibu, seRibu)
}

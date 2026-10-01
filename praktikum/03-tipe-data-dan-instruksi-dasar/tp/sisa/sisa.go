package main

import "fmt"

func main() {
	
	var x, y int
	fmt.Println("Jumlah Kue")
	fmt.Scan(&y)
	fmt.Println("Jumlah Anggota Keluarga")
	fmt.Scan(&x)

	modulo := y % x

	fmt.Printf("Sisa = %d", modulo)
}
package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	tambah := a + b
	kurang := a - b
	kali := a * b
	bagi := a / b
	modulo := a % b

	fmt.Printf("Tambah = %d, Kurang = %d, Kali = %d, Bagi = %d, Modulo = %d\n", tambah, kurang, kali, bagi, modulo)
}

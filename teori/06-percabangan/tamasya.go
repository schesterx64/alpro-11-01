package main
import "fmt"

func main() {
	var n, mobil uint8
	fmt.Scan(&n)
	mobil = n / 7

	if n % 7 != 0 {
		mobil++
	}
	fmt.Println("Jumlah mobil yang diperlukan:", mobil)
}
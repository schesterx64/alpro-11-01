package main
import "fmt"

func main() {
	lamp_status := "Kuning"

	if lamp_status == "Merah" {
		fmt.Println("Berhenti")
	} else if lamp_status == "Kuning" {
		fmt.Println("Hati-Hati / Lebih Pelan")
	} else if lamp_status == "Hijau" {
		fmt.Println("Jalan")
	} else {
		fmt.Println("Perintah tidak terdaftar disini")
	}
}
package main
import "fmt"

func main() {
	var usia int
	var kelas, hKeberangkatan string
	hargaDasar := 0.0
	harga := 0.0
	diskon := 0.0
	fmt.Scan(&kelas, &usia, &hKeberangkatan)

	switch kelas {
	case "Ekonomi":
		hargaDasar = 100000
	case "Bisnis":
		hargaDasar = 200000
	case "Eksekutif":
		hargaDasar = 350000
	default:
		fmt.Println("Kelas tidak valid.")
		return
	}
	if hKeberangkatan == "Sabtu" || hKeberangkatan == "Minggu" {
		hargaDasar *= 1.20
	}
	harga = hargaDasar
	if usia < 3 {
		harga = 0
	} else if usia >= 3 && usia <= 12 {
		diskon = harga * 0.50
		harga -= diskon
	} else if usia >= 60 {
		diskon = harga * 0.30
		harga -= diskon
	}
	fmt.Printf("%.0f\n", harga)
}
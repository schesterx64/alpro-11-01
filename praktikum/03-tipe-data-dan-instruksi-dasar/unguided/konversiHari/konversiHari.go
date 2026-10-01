package main

import "fmt"

func main() {
	var hari, sisaHari, minggu, bulan, tahun int
	fmt.Println("Masukkan jumlah hari:")
	fmt.Scan(&hari)

	tahun = hari / 360
	hari = hari % 360
	bulan = hari / 30
	hari = hari % 30
	minggu = hari / 7
	sisaHari = hari % 7
	fmt.Printf("%d tahun, %d bulan, %d minggu, %d hari\n", tahun, bulan, minggu, sisaHari)
}
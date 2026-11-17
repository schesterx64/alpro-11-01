package main
import "fmt"
const NMAX int = 1001
type tabMurid [NMAX] string

func cariMurid(T tabMurid, n int, nama string, posisi *int, ketemu *bool) {
	for i := 0; i < n; i++ {
		if T[i] == nama {
			*ketemu = true
			*posisi = i + 1
			return
		}
	}
}
func hitungHuruf(T tabMurid, n int, h byte, jumlah *int) {
	for i := 0; i < n; i++ {
		if len(T[i]) > 0 && T[i][0] == h {
			*jumlah++
		}
	}
}
func main() {
	var data tabMurid
	var n, posisi, jumlah int
	var cari string
	var ketemu bool

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}
	fmt.Scan(&cari)
	cariMurid(data, n, cari, &posisi, &ketemu)

	if ketemu {
		fmt.Printf("Murid terdaftar dan berada di urutan absen ke-%d\n", posisi)
	} else {
		fmt.Println("Murid tidak terdaftar")
	}

	var awal byte = cari[0]
	hitungHuruf(data, n, awal, &jumlah)
	fmt.Printf("Jumlah murid dengan huruf awal %c: %d\n", awal, jumlah)
}
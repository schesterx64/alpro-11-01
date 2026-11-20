package main
import "fmt"
const NMAX int = 2002
type tabNIM [NMAX]int

func cariNIM(T tabNIM, n, x int, idx *int) {
	*idx = -1
	for i := 0; i < n && *idx == -1; i++ {
		if T[i] == x {
			*idx = i
		}
	}
}
func ubahNIM(T *tabNIM, idx, baru int) {
	T[idx] = baru
}
func cetak(T tabNIM, n int) {
	for i := 0; i < n; i++ {
		fmt.Println(T[i])
	}
}
func main() {
	var n, cari, pengganti, idx int
	var daftarNIM tabNIM
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&daftarNIM[i])
	}
	
	fmt.Scan(&cari)
	cariNIM(daftarNIM, n, cari, &idx)

	if idx != -1 {
		fmt.Scan(&pengganti)
		ubahNIM(&daftarNIM, idx, pengganti)
		fmt.Printf("Mahasiswa dengan NIM %d terdaftar pada posisi ke-%d\n", cari, idx + 1)
		fmt.Println("Daftar NIM setelah terjadi perubahan data:")
		cetak(daftarNIM, n)
	} else {
		fmt.Println("NIM tidak ditemukan dalam sistem")
	}
}
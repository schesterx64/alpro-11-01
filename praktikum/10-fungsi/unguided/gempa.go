package main
import "fmt"
const NMAX int = 999
type gempa struct {
	nama string
	tahun int
	magnitudo float64
}
type tabGempa [NMAX]gempa

func kategori(m float64) string {
	if m < 4.0 {
		return "Minor"
	} else if m >= 4.0 && m < 6.0 {
		return "Sedang"
	} else if m >= 6.0 && m < 7.0 {
		return "Kuat"
	} else {
		return "Besar"
	}
}
func rataMagnitudo(T tabGempa, n int) float64 {
	var total float64
	for i := 0; i < n; i++ {
		total = total + T[i].magnitudo
	}
	return total / float64(n)
}
func idxTerkuat(T tabGempa, n int) int {
	var idxMax int
	for i := 0; i < n; i++ {
		if T[i].magnitudo > T[idxMax].magnitudo {
			idxMax = i
		}
	}
	return idxMax
}
func hitungKategori(T tabGempa, n int, k string) int {
	var hitung int
	for i := 0; i < n; i++ {
		if kategori(T[i].magnitudo) == k {
			hitung += 1
		}
	}
	return hitung
}
func main() {
	var n, idxMax, Besar int
	var rataRata float64
	var dataGempa tabGempa
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&dataGempa[i].nama, &dataGempa[i].tahun, &dataGempa[i].magnitudo)
	}
	for i := 0; i < n; i++ {
		fmt.Printf("%s %d %.1f %s\n", dataGempa[i].nama, dataGempa[i].tahun, dataGempa[i].magnitudo, kategori(dataGempa[i].magnitudo))
	}
	rataRata = rataMagnitudo(dataGempa, n)
	idxMax = idxTerkuat(dataGempa, n)
	Besar = hitungKategori(dataGempa, n, "Besar")

	fmt.Printf("Rata-rata magnitudo: %.2f\n", rataRata)
	fmt.Printf("Gempa terkuat: %s %d %.1f\n", dataGempa[idxMax].nama, dataGempa[idxMax].tahun, dataGempa[idxMax].magnitudo)
	fmt.Println("Jumlah gempa kategori Besar:", Besar)
}
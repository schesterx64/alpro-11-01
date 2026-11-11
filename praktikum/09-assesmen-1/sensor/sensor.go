package main
import "fmt"

func main() {
	var kelembapan, tertinggi, terendah, total, tKelembapan, sebelumnya int
	var rataRata float64
	fmt.Scan(&kelembapan)
	tertinggi = kelembapan
	terendah = kelembapan
	total = kelembapan
	tKelembapan = 1
	sebelumnya = kelembapan

	for {
		fmt.Scan(&kelembapan)
		if kelembapan == sebelumnya {
			break
		}
		total += kelembapan
		tKelembapan++
		if kelembapan > tertinggi {
			tertinggi = kelembapan
		}
		if kelembapan < terendah {
			terendah = kelembapan
		}
		sebelumnya = kelembapan
	}
	rataRata = float64(total) / float64(tKelembapan)
	fmt.Printf("%d %d %.2f\n", tertinggi, terendah, rataRata)
}
package main
import "fmt"
const NMAX int = 999
type rapot struct {
	matkul string
	nilai int
}
type tabRapot [NMAX]rapot

func nilaiHuruf(x int) string {
	if x > 80 && x <= 100 {
		return "A"
	} else if x > 70 && x <= 80 {
		return "AB"
	} else if x > 65 && x <= 70 {
		return "B"
	} else if x > 60 && x <= 65 {
		return "BC"
	} else if x > 50 && x <= 60 {
		return "C"
	} else if x > 40 && x <= 50 {
		return "D"
	} else {
		return "E"
	}
}
func rataRata(T tabRapot, n int) float64 {
	total := 0
	for i := 1; i <= n; i++ {
		total += T[i].nilai
	}
	return float64(total) / float64(n)
}
func main() {
	var data tabRapot 
	var n int
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Scan(&data[i].matkul, &data[i].nilai)
	}
	for i := 1; i <= n; i++ {
		fmt.Printf("%s %d %s\n", data[i].matkul, data[i].nilai, nilaiHuruf(data[i].nilai))
	}
	fmt.Printf("Rata-rata nilai: %.2f\n", rataRata(data, n))
}
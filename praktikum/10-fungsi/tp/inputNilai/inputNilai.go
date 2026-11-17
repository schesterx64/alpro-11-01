package main
import "fmt"
const NMAX int = 100
type student struct {
	nim, name string
	grade float64
}
type students [NMAX]student

func cariNIM(T students, n int, nim string) int {
	for i := 0; i < n; i++ {
		if T[i].nim == nim {
			return i
		}
	}
	return -1
}
func rataRata(T students, n int) float64 {
	if n == 0 {
		return 0
	}
	var total float64 = 0
	for i := 0; i < n; i++ {
		total += T[i].grade
	}
	return total / float64(n)
}
func main() {
	var data students
	var n int = 0
	var nim, name string
	var grade float64

	for {
		fmt.Scan(&nim)
		if nim == "STOP" {
			break
		}
		fmt.Scan(&name, &grade)

		if cariNIM(data, n, nim) == -1 && n < NMAX {
			data[n].nim = nim
			data[n].name = name
			data[n].grade = grade
			n++
		}
	}
	for i := 0; i < n; i++ {
		fmt.Println(data[n].nim, data[i].name, data[i].grade)
	}
	fmt.Printf("Rata-rata nilai: %.2f\n", rataRata(data, n))
}

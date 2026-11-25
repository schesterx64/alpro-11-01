package main
import "fmt"
const NMAX int = 100
type applicant struct {
	id, name string
	testGrade float64
	testDuration int
}
type applicants [NMAX]applicant

func salinLolos(A applicants, n int, batas float64, B *applicants, m *int) {
	*m = 0
	for i := 0; i < n; i++ {
		if A[i].testGrade >= batas {
			B[*m] = A[i]
			*m++
		}
	}
}
func rataDurasi(B applicants, m int) float64 {
	total := 0
	for i := 0; i < m; i++ {
		total += B[i].testDuration
	}
	return float64(total) / float64(m)
}
func main() {
	var A, B applicants
	var n, m int
	var batas, rataRata float64
	var temp string
	n = 0

	fmt.Scan(&temp)
	for temp != "END" && n < NMAX {
		A[n].id = temp
		fmt.Scan(&A[n].name, &A[n].testGrade, &A[n].testDuration)
		n++
		fmt.Scan(&temp)
	}
	
	fmt.Scan(&batas)
	salinLolos(A, n, batas, &B, &m)

	fmt.Println("Data awal:")
	for i := 0; i < n; i++ {
		fmt.Printf("%s %s %.1f %d\n", A[i].id, A[i].name, A[i].testGrade, A[i].testDuration)
	}

	fmt.Printf("Pelamar lolos (nilai >= %.1f):\n", batas)
	if m == 0 {
		fmt.Println("Tidak ada pelamar yang lolos")
	} else {
		for i := 0; i < m; i++ {
			fmt.Printf("%s %s %.1f %d\n", B[i].id, B[i].name, B[i].testGrade, B[i].testDuration)
		}
		rataRata = rataDurasi(B, m)
		fmt.Println("Jumlah pelamar lolos:", m)
		fmt.Printf("Rata-rata durasi: %.2f menit\n", rataRata)
	}
}

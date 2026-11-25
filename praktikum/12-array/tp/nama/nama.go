package main
import "fmt"
const NMAX int = 9999
type arrString [NMAX]string

func bacaData(A *arrString, n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&A[i])
	}
}
func cetakData(A arrString, n int) {
	for i := 0; i < n; i++ {
		if i < n - 1 {
			fmt.Print(A[i], ", ")
		} else {
			fmt.Println(A[i] + ".")
		}
	}
}
func sisipData(A *arrString, n *int, p int, x string) {
	for i := *n; i >= p; i-- {
		A[i] = A[i-1]
	}
	A[p-1] = x
	*n++
}
func hapusData(A *arrString, n *int, q int) {
	for i := q - 1; i < *n-1; i++ {
		A[i] = A[i+1]
	}
	*n--
}
func main() {
	var A arrString
	var n, p, q int
	var x string

	fmt.Scan(&n)
	bacaData(&A, n)
	fmt.Scan(&p, &x)
	fmt.Scan(&q)

	fmt.Print("Data awal: ")
	cetakData(A, n)

	sisipData(&A, &n, p, x)
	fmt.Print("Setelah penyisipan: ")
	cetakData(A, n)

	hapusData(&A, &n, q)
	fmt.Print("Setelah penghapusan: ")
	cetakData(A, n)
}
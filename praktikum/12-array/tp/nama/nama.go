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
			fmt.Println(A[i], ", ")
		} else {
			fmt.Println(A[i] + ".")
		}
	}
}
func sisipData(A *arrString, n *int, p int, x string) {
	for i := *n; i >= p; 
}
package main
import "fmt"

func hitungLuas(alas, tinggi int, luas *float64) {
	*luas = 0.5 * float64(alas) * float64(tinggi)
}
func main() {
	var a, t, n int
	var luas float64
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a, &t)
		hitungLuas(a, t, &luas)
		fmt.Println(luas)
	}
}
package main
import "fmt"
const NMAX int = 100
type book struct {
	code, title string
	borrowed, pages int
}
type books [NMAX]book

func bacaData(T *books, n *int) {
	var kode string
	*n = 0
	fmt.Scan(&kode)
	for kode != "SELESAI" && *n < NMAX {
		T[*n].code = kode
		fmt.Scan(&T[*n].title, &T[*n].borrowed, &T[*n].pages)
		*n++
		fmt.Scan(&kode)
	}
}
func cetakData(T books, n int) {
	for i := 0; i < n; i++ {
		fmt.Println(T[i].code, T[i].title, T[i].borrowed, T[i].pages)
	}
}
func hapusBuku(T *books, n *int, x int) {
	i := 0
	for i < *n {
		if T[i].borrowed < x {
			for j := i; j < *n-1; j++ {
				T[j] = T[j+1]
			}
			*n--
		} else {
			i++
		}
	}
}
func totalPinjam(T books, n int) int {
	total := 0
	for i := 0; i < n; i++ {
		total += T[i].borrowed
	}
	return total
}
func idxFavorit(T books, n int) int {
	idxMax := 0
	for i := 0; i < n; i++ {
		if T[i].borrowed > T[idxMax].borrowed || (T[i].borrowed ==  T[idxMax].borrowed && T[i].pages < T[idxMax].pages) {
			idxMax = i
		}
	}
	return idxMax
}
func main() {
	var T books
	var n, x int

	bacaData(&T, &n)
	fmt.Scan(&x)

	fmt.Println("Data awal:")
	cetakData(T, n)

	hapusBuku(&T, &n, x)

	fmt.Printf("Data setelah penghapusan (peminjaman >= %d):\n", x)

	if n == 0 {
		fmt.Println("Tidak ada buku tersisa")
	} else {
		cetakData(T, n)
		fmt.Println("Total peminjaman:", totalPinjam(T, n))
		idxFav := idxFavorit(T, n)
		fmt.Println("Buku paling sering dipinjam:")
		fmt.Println(T[idxFav].code, T[idxFav].title, T[idxFav].borrowed, T[idxFav].pages)
	}
}
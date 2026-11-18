package main
import "fmt"
const NMAX int = 1000
type penduduk struct {
	nama string
	tahun int
	kota string
}
type tabPenduduk [NMAX]penduduk

func cariTahun(T tabPenduduk, n, th int, idx *int) {
	*idx = -1
	for i := 0; i < n && *idx == -1; i++ {
		if T[i].tahun == th {
			*idx = i
		}
	}
}
func tertuaTermuda(T tabPenduduk, n int, tua, muda *penduduk) {
	*tua = T[0]
	*muda = T[0]
	for i := 1; i < n; i++ {
		if T[i].tahun < tua.tahun {
			*tua = T[i]
		}
		if T[i].tahun > muda.tahun {
			*muda = T[i]
		}
	}
}
func main() {
	var data tabPenduduk
	var n, thCari, idx int
	var tua, muda penduduk
	
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i].nama, &data[i].tahun, &data[i].kota)
	}

	fmt.Scan(&thCari)
	cariTahun(data, n, thCari, &idx)
	if idx != -1 {
		fmt.Printf("%s %d %s ditemukan di index ke-%d\n", data[idx].nama, data[idx].tahun, data[idx].kota, idx)
	} else {
		fmt.Println("Data tidak ditemukan")
	}

	tertuaTermuda(data, n, &tua, &muda)
	fmt.Printf("Penduduk tertua: %s (%d)\n", tua.nama, tua.tahun)
	fmt.Printf("Penduduk termuda: %s (%d)\n", muda.nama, muda.tahun)
}
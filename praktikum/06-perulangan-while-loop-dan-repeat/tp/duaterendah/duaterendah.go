package main
import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	min1, min2 := 999999, 999999
	namaMin1, namaMin2 := "", ""

	for i := 0; i < n; i++ {
		var nama string
		var nilai int
		fmt.Scan(&nama, &nilai)

		if nilai < min1 {
			min2 = min1
			namaMin2 = namaMin1
			min1 = nilai
			namaMin1 =  nama
		} else if nilai < min2 {
			min2 = nilai
			namaMin2 = nama
		}
	}
	fmt.Println(namaMin1)
	fmt.Println(namaMin2)
}
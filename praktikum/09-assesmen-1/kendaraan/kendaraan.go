package main
import "fmt"

func main() {
	var kendaraan string
	m := 0
	k := 0
	t := 0

	for {
		fmt.Scan(&kendaraan)

		switch kendaraan {
		case "M":
			m++
		case "K":
			k++
		case "T":
			t++
		}

		if kendaraan != "M" && kendaraan != "K" && kendaraan != "T" {
			break
		}

	}
	fmt.Printf("Tipe M = %d, Tipe K = %d, Tipe T = %d", m, k, t)
}
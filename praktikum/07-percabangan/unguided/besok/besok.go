package main
import "fmt"

func main() {
	var tg, bl, th, hari int
	fmt.Scan(&tg, &bl, &th)
	kabisat := (th % 4 == 0 && th % 100 != 0) || (th % 400 == 0)

	switch bl {
	case 1, 3, 5, 7, 8, 10, 12:
		hari = 31
	case 4, 6, 9, 11:
		hari = 30
	case 2:
		if kabisat {
			hari = 29
		} else {
			hari = 28
		}
	}
	if bl >= 1 && bl <= 12 && tg >= 1 && tg <= hari {
		fmt.Println("Valid")

		tg++
		if tg > hari {
			tg = 1
			bl++
			if bl > 12 {
				bl = 1
				th++
			}
		}
		fmt.Printf("%d/%d/%d\n", tg, bl, th)
	} else {
		fmt.Println("Tanggal tidak valid")
	}
}
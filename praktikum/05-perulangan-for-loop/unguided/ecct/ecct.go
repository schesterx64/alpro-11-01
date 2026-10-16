package main
import "fmt"

func main() {
	var m int 
	fmt.Scan(&m)
	lulus := 0

	for i := 1; i <= m; i++ {
		var P, V, GA, FS, CI, TC int
		fmt.Scan(&P, &V, &GA, &FS, &CI,&TC)

		total := P + V + GA + FS + CI + TC
		rataRata := float64(total) / 6.0

		if rataRata >= 3.33 {
			fmt.Println("Ya")
			lulus++
		} else {
			fmt.Println("Tidak")
		}
	}
	fmt.Println(lulus)
}
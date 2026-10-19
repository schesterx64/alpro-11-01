package main
import "fmt"

func main() {
	var total, diskon float64
	var clo bool
	diskon = 0.35
	fmt.Scan(&total, &clo)

	if clo == true {
		totalDiskon  := total * diskon
		fmt.Println(total - totalDiskon)
	} else if clo == false {
		fmt.Println(total)
	}
}
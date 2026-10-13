package main
import "fmt"

func main() {
	var x int
	//var ganjil int = x % 2 != 0
	fmt.Scan(&x)
	//ganjil := x % 2 != 0
	total := 0

	for i := 1; i <= x; i++{
		if i % 2 != 0 {
			fmt.Println(i)
			total += i
		}
	}
	fmt.Println(total)
}
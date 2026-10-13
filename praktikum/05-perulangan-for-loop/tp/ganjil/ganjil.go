package main
import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	total := 0

	for i := 1; i <= x; i++{
		if i % 2 != 0 {
			fmt.Println(i)
			total += i
		}
	}
	fmt.Println(total)
}
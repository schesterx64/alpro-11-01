package main
import "fmt"

func main() {
	var n int = 1
	fmt.Scan(&n)
	for n <= 10 {
		fmt.Println(n)
		n++
	}
}
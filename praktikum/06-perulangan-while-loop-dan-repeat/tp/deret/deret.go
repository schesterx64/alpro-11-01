package main
import "fmt"

func main() {
	var n int
	jumlah := 0
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		jumlah += i
	}
	fmt.Println(jumlah)
}
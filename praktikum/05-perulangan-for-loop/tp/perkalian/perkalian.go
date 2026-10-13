package main
import "fmt"

func main() {
	var a, b int
	fmt.Print("Bilangan ")
	fmt.Scan(&a)
	fmt.Print("Dengan batas perkalian ")
	fmt.Scan(&b)

	fmt.Println("Hasil: ")
	for i := 1; i <= b; i++ {
		fmt.Printf("%d x %d = %d\n", a, i, a * i)
	}
}
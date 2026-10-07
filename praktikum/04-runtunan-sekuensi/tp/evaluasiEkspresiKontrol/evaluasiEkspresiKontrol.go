package main

import "fmt"

func main() {
	intNum := 5
	intOther := 10
	var sngNum float64 = -3

	if intNum > 0 {
		fmt.Println(intNum > 0)
	}

	if intNum >= 5 && intOther < 11 {
		fmt.Println(intNum >= 5 && intOther < 11)
	}

	if sngNum > 0 || (intNum > 0 && -1 * intOther == -10) {
		fmt.Println(sngNum > 0 || (intNum > 0 && -1 * intOther == -10))
	}
}
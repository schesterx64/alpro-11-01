package main

import "fmt"

func main() {
	var name string

	name = "Zeda Briyan Dinillah"

	fmt.Println("Nama =", name)

	var lastName = "Dinillah"
	fmt.Println("Nama Belakang =", lastName)

	middleName := "Briyan"
	fmt.Println("Nama Tengah =", middleName)

	var (
		fullName = "Zeda Briyan Dinillah"
		firstName = "Zeda"
	)

	fmt.Println(fullName)
	fmt.Println(firstName)
}

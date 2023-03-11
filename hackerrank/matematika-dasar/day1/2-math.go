package main

import "fmt"

func main () {
	var a, b, c, hasil int

	fmt.Print("Masukan nilai a: ")
	fmt.Scanln(&a)
	fmt.Print("Masukan nilai b: ")
	fmt.Scanln(&b)
	fmt.Print("Masukan nilai c: ")
	fmt.Scanln(&c)

	hasil = a / b * c

	fmt.Println("Hasilnya:")
	fmt.Printf("%d : %d x %d = %d", a, b, c, hasil)
}
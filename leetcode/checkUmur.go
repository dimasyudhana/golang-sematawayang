package main

import "fmt"

func main() {
	var firstname string
	var lastname string
	var umur int

	fmt.Print("Masukkan nama depan: ")
	fmt.Scan(&firstname)

	fmt.Print("Masukkan nama belakang: ")
	fmt.Scan(&lastname)

	fmt.Print("Masukkan umur: ")
	fmt.Scan(&umur)

	fmt.Printf("Hi %s %s, umurmu %d, maka ", firstname, lastname, umur)

	if umur <= 15 {
		fmt.Print("muda")
	} else if umur <= 25 {
		fmt.Print("remaja")
	} else if umur > 25 && umur <= 50 {
		fmt.Print("dewasa")
	} else {
		fmt.Print("pensiun")
	}
}

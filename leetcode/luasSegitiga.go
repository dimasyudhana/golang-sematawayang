package main

import "fmt"

func main() {
	var alas float64
	fmt.Print("Input alas : ")
	fmt.Scan(&alas)

	var tinggi float64
	fmt.Print("Input tinggi : ")
	fmt.Scan(&tinggi)

	var luasSegitiga float64 = (alas * tinggi) / 2
	fmt.Printf("luas segitiga dari (%g * %g) / 2 = %g", alas, tinggi, luasSegitiga)
}

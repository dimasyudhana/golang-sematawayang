package main

import (
	"fmt"
	"math"
)

func main (){
	var a,b,c,d float64
	var hasil float64

	fmt.Print("Masukan nilai a: ")
	fmt.Scanln(&a)
	fmt.Print("Masukan nilai b: ")
	fmt.Scanln(&b)
	fmt.Print("Masukan nilai c: ")
	fmt.Scanln(&c)
	fmt.Print("Masukan nilai d: ")
	fmt.Scanln(&d)

	hasil = math.Pow(a - b, 2) + math.Pow(c, 2) / math.Cbrt(d)

	fmt.Println("Hasilnya :")
	fmt.Printf("(%.f - %.f)² + %.f² / ∛%.f = %.f", a, b, c, d, hasil)
}
package main

import (
	"fmt"
	"math"
)

func main() {
	var num int
	fmt.Print("Masukkan sebuah bilangan: ")
	fmt.Scan(&num)

	if isKeren(num) {
		fmt.Printf("%d adalah bilangan keren.\n", num)
	} else {
		fmt.Printf("%d bukan bilangan keren.\n", num)
	}
}

func isKeren(num int) bool {
	if num <= 1 {
		return false
	}

	// Memeriksa apakah num habis dibagi oleh 1 dan dirinya sendiri
	if num%1 != 0 || num%(num) != 0 {
		return false
	}

	// Mencari faktor dari num, kecuali 1 dan num itu sendiri
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			// Memeriksa apakah faktor tersebut hanya bisa dibagi oleh 1 dan dirinya sendiri
			if !isPrime(i) || !isPrime(num/i) {
				return false
			}
		}
	}

	return true
}

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}

	// Memeriksa apakah num habis dibagi oleh bilangan selain 1 dan dirinya sendiri
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

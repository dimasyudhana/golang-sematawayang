package main

import "fmt"

func main() {
	var N int
	fmt.Print("Input angka : ")
	fmt.Scan(&N)
	//check ganjil - genap
	if N%2 == 0 {
		fmt.Println("Genap")
	} else {
		fmt.Println("Ganjil")
	}
	return
}

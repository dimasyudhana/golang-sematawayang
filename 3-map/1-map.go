package main

import "fmt"

func main() {
	var chicken map[string]int //ada error
	chicken = map[string]int{} //tidak ada error

	chicken["Januari"] = 30
	chicken["Februari"] = 40
	chicken["Maret"] = 50
	chicken["April"] = 60
	chicken["Mei"] = 35

	fmt.Println("Januari", chicken["Januari"]) //Januari 30
	fmt.Println("Maret", chicken["Maret"])     //Maret 50
	fmt.Println("Juni", chicken["Juni"])       //Maret 50
}

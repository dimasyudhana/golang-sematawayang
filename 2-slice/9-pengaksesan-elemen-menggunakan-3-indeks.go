package main

import "fmt"

func main() {
	var fruits = []string{"apple", "grape", "banana", "apple", "grape", "banana"}
	var aFruits = fruits[0:2]
	var bFruits = fruits[0:2:2]
	var cFruits = fruits[2:4:6] //merepresentasikan kapasitas yang di akses walaupun yang ditampilkan hanya 2 elemen.

	fmt.Println(fruits)      // ["apple", "grape", "banana"]
	fmt.Println(len(fruits)) // len: 3
	fmt.Println(cap(fruits)) // cap: 3

	fmt.Println(aFruits)      // ["apple", "grape"]
	fmt.Println(len(aFruits)) // len: 2
	fmt.Println(cap(aFruits)) // cap: 3

	fmt.Println(bFruits)      // ["apple", "grape"]
	fmt.Println(len(bFruits)) // len: 2
	fmt.Println(cap(bFruits)) // cap: 2

	fmt.Println(cFruits)      // ["apple", "grape"]
	fmt.Println(len(cFruits)) // len: 2
	fmt.Println(cap(cFruits)) // cap: 4
}

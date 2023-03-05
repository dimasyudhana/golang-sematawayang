package main

import "fmt"

func main() {
	var fruits = [8]string{
		"apple",
		"grape",
		"banana",
		"melon",
		"papaya",
		"orange",
		"kiwi",
		"avocado",
	}
	for i, fruit := range fruits {
		fmt.Printf("elemen %d: fruit_name: %s\n", i, fruit)
	}
}

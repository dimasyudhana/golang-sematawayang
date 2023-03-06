package main

import "fmt"

func main() {
	var chicken = map[string]int{
		"Januari":  30,
		"Februari": 40,
		"Maret":    50,
		"April":    60,
		"Mei":      35,
	}

	fmt.Println(len(chicken))
	fmt.Println(chicken)

	delete(chicken, "Januari")

	fmt.Println(len(chicken))
	fmt.Println(chicken)
}

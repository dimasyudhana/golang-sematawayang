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
	var value, isExist = chicken["Juni"]

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item is not exists")
	}
}

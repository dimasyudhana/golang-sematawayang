package main

import "fmt"

func main() {
	var chicken1 = map[string]int{"Januari": 30, "Februari": 40, "Maret": 50}

	var chicken2 = map[string]int{
		"Januari":  30,
		"Februari": 40,
		"Maret":    50,
	}
	var chicken3 = map[string]int{}
	var chicken4 = make(map[string]int)
	var chicken5 = *new(map[string]int)

	fmt.Println(chicken1) //map[Februari:40 Januari:30 Maret:50]
	fmt.Println(chicken2) //map[Februari:40 Januari:30 Maret:50]
	fmt.Println(chicken3) //map[]
	fmt.Println(chicken4) //map[]
	fmt.Println(chicken5) //map[]
}

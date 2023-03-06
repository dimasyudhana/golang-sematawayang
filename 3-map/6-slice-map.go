package main

import "fmt"

func main() {
	var chickens = []map[string]string{
		{"name": "chicken blue", "gender": "male"},
		{"name": "chicken red", "gender": "female"},
		{"name": "chicken green", "gender": "male"},
	}
	for _, chicken := range chickens {
		fmt.Println(chicken["gender"], chicken["name"])
	}
}

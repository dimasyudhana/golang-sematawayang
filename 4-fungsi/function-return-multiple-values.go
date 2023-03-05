package main

import "fmt"

func getFullName() (string, string) {
	return "Dimas", "Yudhana"
}

func main() {
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)
}

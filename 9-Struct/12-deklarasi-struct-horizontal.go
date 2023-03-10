package main

import "fmt"

func main() {
	type person struct {
		name    string
		age     int
		hobbies []string
	}

	var p1 = struct {
		name string
		age  int
	}{age: 22, name: "Shelby"}
	var p2 = struct {
		name string
		age  int
	}{"Boast", 25}

	fmt.Println(p1) // {Shelby 22}
	fmt.Println(p2) // {Boast 25}
}
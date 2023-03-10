package main

import "fmt"

type Person struct {
	name string `tag1`
	age int `tag2`
}

type People = Person

func main () {
	var p1 = Person{"Shelby", 21}
	fmt.Println(p1) // {Shelby 21}

	var p2 = People{"Tommy", 25}
	fmt.Println(Person(p1)) // {Shelby 21}
	fmt.Println(Person(p2)) // {Tommy 25}
}
package main

import "fmt"

type person struct {
	name string
	age int
}

type student struct {
	person
	age int
	grade int
}

func main () {
	var s1 = student{}
	s1.name = "tommy"
	s1.age = 21		//age og student
	s1.person.age = 22 //age of person

	fmt.Println(s1.name) // tommy
	fmt.Println(s1.age) // 21
	fmt.Println(s1.person.age) //22
}
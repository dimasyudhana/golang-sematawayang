package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	var s1 = struct {
		person
		grade int
	}{}

	s1.person = person{"shelby", 21}
	s1.grade = 2

	fmt.Println("name	:", s1.person.name)
	fmt.Println("age	:", s1.person.age)
	fmt.Println("grade	:", s1.grade)

	var s2 = struct {
		person
		grade int
	}{}

	s2.person = person{"Tommy", 21}
	s2.grade = 2

	fmt.Println("name	:", s2.person.name)
	fmt.Println("age	:", s2.person.age)
	fmt.Println("grade	:", s2.grade)
}
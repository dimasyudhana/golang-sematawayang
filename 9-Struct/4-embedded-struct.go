package main

import "fmt"

type person struct {
	name string
	age  int
}

type student struct {
	grade int
	person
}

func main() {
	var s1 = student{}
	s1.name = "tommy"
	s1.age = 21
	s1.grade = 2

	fmt.Println("name : ",s1.name) // name :  tommy
	fmt.Println("age : ",s1.age) // age :  21
	fmt.Println("age : ",s1.person.age) // age :  21
	fmt.Println("grade : ",s1.grade) // grade :  2
}
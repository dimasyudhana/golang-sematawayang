package main

import "fmt"

type person struct {
	name string
	age int
}

func main(){
	var student struct {
		person
		grade int
	}

	student.person = person{"Shelby", 21}
	student.grade = 2

	fmt.Println(student.name)	// Shelby
	fmt.Println(student.age)	// 21
	fmt.Println(student.grade)	// 2
}
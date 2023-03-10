package main

import "fmt"

type person struct {
	name string
	age int
}

func main () {
	var allStudents = []person{
		{name: "Thomas", age: 23},
		{name: "Shelby", age: 24},
		{name: "Boaz", age: 25},
	}

	for _, student := range allStudents {
		fmt.Println(student.name, "age is", student.age)
	}
}
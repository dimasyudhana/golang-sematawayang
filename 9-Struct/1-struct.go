package main

import "fmt"

type student struct {
	name string
	grade int
}

func main() {
	var s1 student
	s1.name = "john wick"
	s1.grade = 2

	var s2 student
	s2.name = "dimas yudhana"
	s2.grade = 3

	fmt.Println("name : ", s1.name)
	fmt.Println("grade : ", s1.grade)
	
	fmt.Println("name : ", s2.name)
	fmt.Println("grade : ", s2.grade)
}
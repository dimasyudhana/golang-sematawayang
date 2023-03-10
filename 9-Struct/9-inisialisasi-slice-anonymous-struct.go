package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	var allStudents = []struct {
		person
		grade int
	}{
		{person: person{"thomas", 22}, grade: 2},
		{person: person{"shelby", 23}, grade: 3},
		{person: person{"boaz", 24}, grade: 3},
	}

	for _, student := range allStudents {
		fmt.Println(student)
	}
}
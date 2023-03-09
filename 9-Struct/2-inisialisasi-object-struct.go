package main

import (
	"fmt"
)

type student struct {
	name  string
	grade int
}

func main() {
	var s1 = student{}
	s1.name = "thomas"
	s1.grade = 2

	var s2 = student{"ethan", 2}

	var s3 = student{name : "made"}

	fmt.Println("student 1 : ", s1.name)
	fmt.Println("student 1 : ", s2.name)
	fmt.Println("student 1 : ", s3.name)
}
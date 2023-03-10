package main

import "fmt"

type student struct {
	name string
	grade int
}

func main () {
	var s1 = student {
		name : "tommy", 
		grade: 2,
	}

	var s2 *student = &s1
	
	fmt.Println("student 1, name : ", s1.name) //student 1, name :  tommy
	fmt.Println("student 2, name : ", s2.name) //student 2, name :  tommy

	s2.name = "shelby"

	fmt.Println("student 1, name : ", s1.name) // student 1, name :  shelby
	fmt.Println("student 2, name : ", s2.name) //student 2, name :  shelby
}
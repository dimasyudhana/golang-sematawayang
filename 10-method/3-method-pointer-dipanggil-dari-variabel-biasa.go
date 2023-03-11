package main

import "fmt"

type student struct {
	name string
	grade int
}

func (s student) sayHello() {
	fmt.Println("Hallo", s.name)
}

func main () {
	var s1 = student{"Thomas Shelby", 25}
	s1.sayHello()

	var s2 = &student{"Richard Barker", 25}
	s2.sayHello()
}
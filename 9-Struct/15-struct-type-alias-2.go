package main

import "fmt"

type People1 struct {
	name string
	age int
}

type People2 = struct {
	name string
	age int
}

func main () {
	people := People1{"Shelby", 21}
	fmt.Println(People1(people)) // {Shelby 21}

	person := People2{"Boast", 25}
	fmt.Println(People2(person)) // {Boast 25}
}
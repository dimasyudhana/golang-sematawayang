package main

import "fmt"

func sayHelloTo(firstname string, lastname string) {
	fmt.Println("Hallo ", firstname, "", lastname)
}

func main() {
	firstname := "Angela"
	sayHelloTo(firstname, "Yu")
	for i := 0; i < 3; i++ {
		sayHelloTo("Dimas", "Yudhana")
	}
	sayHelloTo("Made", "Trisna")
}

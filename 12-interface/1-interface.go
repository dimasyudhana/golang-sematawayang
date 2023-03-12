package main

import "fmt"

type HasName interface {
	GetName() string
}

func sayHello(hasName HasName){
	fmt.Println("Hello", hasName.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

//contoh lain walaupun struct berbeda, kita dapat menggunakan kontrak interface yang sama

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main () {
	/*
	// var Terry HasName
	// // Terry.GetName()
	// sayHello(Terry) //panic: runtime error: invalid memory address or nil pointer dereference

	//bagaimana cara implementasi interface dalam main?
	// dengan cara mempuat struct untuk interface tersebut lalu membuat func untuk struct tsb.
	// maka secara otomatis struct akan megikuti kontrak pada interface (HasName)
	*/

	var terry Person
	terry.Name = "Terry"
	
	sayHello(terry)

	cat := Animal {
		Name: "Fluffy",
	}

	sayHello(cat)


}
package main

import (
	"fmt"
	"strings"
)

func yourHobbies(nama string, hobby ...string) {
	var hobbiesAsString = strings.Join(hobby, ", ")

	fmt.Printf("Hello my name is %s\n", nama)
	fmt.Printf("my hobbies are %s\n", hobbiesAsString)

}

func main() {
	yourHobbies("wick", "sleeping", "eating")

	fmt.Println("----------------------------")

	var hobbies = []string{"sleeping", "eating"}
	yourHobbies("wick", hobbies...)
}

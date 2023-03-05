package main

import "fmt"

func sayHello() {
	fmt.Println("Hallo World")
}

// func main() {
// 	sayHello()
// }

func main() {
	for i := 0; i < 10; i++ {
		sayHello()
	}
}

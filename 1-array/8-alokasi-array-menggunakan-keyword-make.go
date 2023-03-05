package main

import "fmt"

func main() {
	var fruits = make([]string, 5)

	fruits[0] = "apple"
	fruits[1] = "banana"
	fruits[2] = "raspberry"
	fruits[3] = "blackberry"
	fruits[4] = "cherry"

	fmt.Println(fruits)
}

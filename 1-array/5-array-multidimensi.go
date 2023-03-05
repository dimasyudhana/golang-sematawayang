package main

import "fmt"

func main() {
	var array_number1 = [2][3]int{[3]int{1, 2, 3}, [3]int{3, 4, 5}}
	var array_number2 = [2][3]int{{1, 2, 3}, {3, 4, 5}}

	fmt.Println("array_number1", array_number1)
	fmt.Println("array_number2", array_number2)
}

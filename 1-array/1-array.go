package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Dimas"
	names[1] = "Ady"
	names[2] = "Yudhana"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
	fmt.Println(names[0], names[1], names[2])
	fmt.Println(names)

	var values = [3]int{
		95,
		97,
		99,
	}

	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])
	fmt.Println(values)

	fmt.Println(len(names))
	fmt.Println(len(values))

	var lainLagi = [100]int{}

	fmt.Println(len(lainLagi))
}

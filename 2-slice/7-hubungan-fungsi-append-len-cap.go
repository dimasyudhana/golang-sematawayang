package main

import "fmt"

func main() {
	var fruits = []string{
		"apple",
		"grape",
		"manggo",
		"cherry",
		"avocado",
	}
	var bfruits = fruits[2:4]

	fmt.Println(fruits)
	fmt.Println(len(fruits))
	fmt.Println(cap(fruits))
	fmt.Println("\n")

	fmt.Println(bfruits)
	fmt.Println(len(bfruits))
	fmt.Println(cap(bfruits))

	var cfruits = append(bfruits, "orange")
	fmt.Println(cfruits)
	fmt.Println(len(cfruits))
	fmt.Println(cap(cfruits))

	var dfruits = append(bfruits, "orange", "orange")
	fmt.Println(dfruits)
	fmt.Println(len(dfruits))
	fmt.Println(cap(dfruits))

	var efruits = append(bfruits, "orange", "orange", "orange", "orange", "orange", "orange", "orange", "orange", "orange", "orange", "orange", "orange", "orange")
	fmt.Println(efruits)
	fmt.Println(len(efruits))
	fmt.Println(cap(efruits)) /// 15 dan 15???
}

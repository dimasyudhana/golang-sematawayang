package main

import "fmt"

func findMax(numbers []int, max int) (int, func() []int) {
	var result []int
	for _, e := range numbers {
		if e <= max {
			result = append(result, e)
		}
	}
	return len(result), func() []int {
		return result
	}
}

func main() {
	var max = 3
	var numbers = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}
	var howMany, getNumber = findMax(numbers, max)
	var theNumbers = getNumber()

	fmt.Println("numbers\t:", numbers)
	fmt.Printf("find\t: %d\n\n", max)
	fmt.Println("found\t:", howMany)
	fmt.Println("value\t:", theNumbers)
}

package main

import "fmt"

func main() {
	devideNumber(10, 5)
	devideNumber(4, 0) //panic : runtime error: integer divide by zero
	devideNumber(5, 2)
}

func devideNumber(a, b int) {
	if b == 0 {
		fmt.Printf("invalid devider. %d cannot devided %d\n", a, b)
		return
	}
	var result = a / b
	fmt.Printf("%d/%d = %d\n", a, b, result)
}

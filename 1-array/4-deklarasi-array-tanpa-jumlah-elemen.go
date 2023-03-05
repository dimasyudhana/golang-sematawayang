package main

import "fmt"

func main() {
	var product_id = [...]int{
		1,
		2,
		3,
		4,
		5,
	}
	fmt.Println("data array \t\t : ", product_id)
	fmt.Println("length data array \t : ", len(product_id))
}

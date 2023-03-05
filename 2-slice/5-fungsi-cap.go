package main

import "fmt"

func main() {
	var product_names = []string{
		"Indomie",
		"Pop Mie",
		"Sarimi",
		"Supermie",
		"Indomilk",
		"Chitato",
		"Wang's",
		"Cap Panda",
		"Cap Kaki Tiga",
		"Granulab",
	}
	fmt.Println(len(product_names))
	fmt.Println(cap(product_names))

	var aproduct_names = product_names[0:5]
	fmt.Println(len(aproduct_names))
	fmt.Println(cap(aproduct_names))

	var bproduct_names = product_names[7:9]
	fmt.Println(len(bproduct_names))
	fmt.Println(cap(bproduct_names))
}

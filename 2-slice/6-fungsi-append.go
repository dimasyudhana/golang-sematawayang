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
	fmt.Println(product_names)
	fmt.Println(len(product_names))
	fmt.Println(cap(product_names))

	var cproduct_names = append(product_names, "Wonglokat")
	fmt.Println(cproduct_names)
	fmt.Println(len(cproduct_names))
	fmt.Println(cap(cproduct_names)) //kapasitas menigkat menjadi 2 kali lipat.
}

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
	var aproduct_names = product_names[0:5]
	var bproduct_names = product_names[5:10]

	var aaproduct_names = aproduct_names[0:5]
	var baproduct_names = bproduct_names[0:5]

	fmt.Println(product_names)   //[Indomie Pop Mie Sarimi Supermie Indomilk Chitato Wang's Cap Panda Cap Kaki Tiga Granulab]
	fmt.Println(aproduct_names)  //[Indomie Pop Mie Sarimi Supermie Indomilk]
	fmt.Println(bproduct_names)  //[Chitato Wang's Cap Panda Cap Kaki Tiga Granulab]
	fmt.Println(aaproduct_names) //[Indomie Pop Mie Sarimi Supermie Indomilk]
	fmt.Println(baproduct_names) //[Chitato Wang's Cap Panda Cap Kaki Tiga Granulab]
}

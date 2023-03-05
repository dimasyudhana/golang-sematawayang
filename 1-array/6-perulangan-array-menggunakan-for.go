package main

import "fmt"

func main() {
	var product_names = [5]string{
		"Indomie",
		"Pop Mie",
		"Mie Telur 3 Ayam",
		"Cap Kaki Tiga",
		"Cap Panda",
	}
	for i := 0; i < len(product_names); i++ {
		fmt.Printf("No. %d\tproduct_name: %s\n", i, product_names[i])
	}
}

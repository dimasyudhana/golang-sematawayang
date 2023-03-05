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
	fmt.Println(product_names)           //[Indomie Pop Mie Sarimi Supermie Indomilk Chitato Wang's Cap Panda Cap Kaki Tiga Granulab]
	fmt.Println(product_names[0:0])      //[]
	fmt.Println(product_names[10:10])    //[]
	fmt.Println(product_names[0:2])      //[Indomie Pop Mie]
	fmt.Println(product_names[3:5])      //[Supermie Indomilk]
	fmt.Println(len(product_names[0:5])) //5
	fmt.Println(product_names[:])        //[Indomie Pop Mie Sarimi Supermie Indomilk Chitato Wang's Cap Panda Cap Kaki Tiga Granulab]
	fmt.Println(product_names[2:])       //[Sarimi Supermie Indomilk Chitato Wang's Cap Panda Cap Kaki Tiga Granulab]
	fmt.Println(product_names[:8])       //[Indomie Pop Mie Sarimi Supermie Indomilk Chitato Wang's Cap Panda]
}

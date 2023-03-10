package main

import "fmt"

type person struct {
	//struct field tag `tag1` not compatible with reflect.StructTag.Get: bad syntax for struct tag pairstructtag
	name string `tag1`
	age int `tag2`
}

func main () {
	var p1 = person{"Shelby", 21}
	fmt.Println(p1) // {Shelby 21}
}
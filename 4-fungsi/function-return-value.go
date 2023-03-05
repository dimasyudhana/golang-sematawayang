package main

import "fmt"

func getHello(nama string) string {
	if nama == "" {
		return "Hello bro"
	} else {
		return "Hello " + nama
	}
}

func main() {
	result := getHello("Angela")
	fmt.Println(result)

	fmt.Println(getHello(""))
}

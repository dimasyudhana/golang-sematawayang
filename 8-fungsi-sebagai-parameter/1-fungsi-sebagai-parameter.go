package main

import (
	"fmt"
	"strings"
)

func filter(data []string, callback func(string) bool) []string {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result
}

func main() {
	var data = []string{"wick", "jason", "ethan"}
	var dataContainsO = filter(data, func(each string) bool {
		return strings.Contains(each, "o")
	})

	var dataLength5 = filter(data, func(each string) bool {
		return len(each) == 5
	})

	//data asli : ["wick", "jason", "ethan"]
	fmt.Println("data asli\t\t:", data)

	//filter ada huruf "o"    : [jason]
	fmt.Println("filter ada huruf \"o\"\t:", dataContainsO)

	//filter jumlah huruf "5" : [jason ethan]
	fmt.Println("filter jumlah huruf \"5\"\t:", dataLength5)
}

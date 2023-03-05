package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 { //kalau 1 maka genap
			continue
		}
		fmt.Println("perulangan ke: ", i)
	}
}

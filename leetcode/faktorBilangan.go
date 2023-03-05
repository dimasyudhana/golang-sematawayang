package main

import "fmt"

func main() {
	var N int
	fmt.Print("faktor bilangan = ")
	fmt.Scan(&N)
	fmt.Printf("faktor bilangan %d = ", N)

	for i := 1; i <= N; i++ {
		if N%i == 0 {
			fmt.Print(i, " ")
		}
	}
}

package main

import (
	"fmt"
)

func FizzBuzz (bilangan int) string {	
	if bilangan%5 == 0 && bilangan%3 == 0{
		return "FizzBuzz"
	}else if bilangan%5 == 0 {
		return "Buzz"	
	}else if bilangan%3 == 0{
		return "Fizz"
	}else{
		return fmt.Sprintf("%d", bilangan)
	}
}

func main () {
	var panjangBilangan int
	fmt.Println("Berapa banyak angka?")
	fmt.Print("Banyak angka: ")
	fmt.Scanln(&panjangBilangan)

	for i := 0; i <= panjangBilangan; i++ {
		fmt.Println(FizzBuzz(i))
	}
}

// 1
// 2
// Fizz
// 4
// Buzz
// Fizz
// 7
// 8
// Fizz
// Buzz
// 11
// Fizz
// 13
// 14
// FizzBuzz
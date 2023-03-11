package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	num := x
	reverse := 0
	for num > 0 {
		reverse = reverse*10 + num%10
		num = num / 10
	}
	return reverse == x
}

func main() {
	fmt.Println(isPalindrome(121))
}

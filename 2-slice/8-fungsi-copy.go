package main

import "fmt"

func main() {
	dst := make([]string, 3)
	src := []string{"watermelon", "apple", "orange", "pinnaple"}
	n := copy(dst, src)

	fmt.Println(dst)
	fmt.Println(src)
	fmt.Println(n)

	// ---------------------

	dst1 := []string{"potato", "potato", "potato"}
	src1 := []string{"watermelon", "pinnaple"}
	n1 := copy(dst1, src1)

	fmt.Println(dst1) // watermelon pinnaple potato
	fmt.Println(src1) // watermelon pinnaple
	fmt.Println(n1)   // 2
}

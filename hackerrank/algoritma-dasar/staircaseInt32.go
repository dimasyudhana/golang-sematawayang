package main

import (
	"fmt"
)

func Staircase(n int32) {
	outputStair := ""
	for i := int32(0); i < n; i++ {
		for j := int32(0); j < n-i-1; j++ {
			outputStair += " "
		}
		for k := int32(0); k <= i; k++ {
			outputStair += "#"
		}
		outputStair += "\n"
	}
	fmt.Println(outputStair)
}

func main() {
	Staircase(4)
}
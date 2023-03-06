package main

import (
	"fmt"
	"math"
)

func calculate(d float64) (float64, float64) {
	//hitung luas
	var area = math.Pi * math.Pow(d/2, 2)
	//hitung keliling
	var circumference = math.Pi * d

	//kembalikan 2 nilai
	return area, circumference
}

func main() {
	var diameter float64 = 15
	var area, circumference = calculate(diameter)

	fmt.Printf("luas lingkaran \t\t: %.2f \n", area) // mencetak 2 digit angka dibelakang koma
	fmt.Printf("keliling lingkaran \t: %.2f \n", circumference)

}

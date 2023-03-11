package main

import "fmt"

// func operasi ()

func main(){
	var a,b,c,hasil int
	
	//input nilai a
	fmt.Print("Masukan nilai a : ")
	fmt.Scanln(&a)

	//input nilai b
	fmt.Print("Masukan nilai b : ")
	fmt.Scanln(&b)

	//input nilai c
	fmt.Print("Masukan nilai c : ")
	fmt.Scanln(&c)

	//hasil operasi matematika
	hasil = a - b + c
	fmt.Println("Hasil dari :")
	fmt.Printf("%d - %d + %d = %d\n",a, b, c, hasil)
}
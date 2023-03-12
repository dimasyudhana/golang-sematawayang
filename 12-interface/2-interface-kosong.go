/*
//Golang bukanlah 	bahasa pemrgraman yang berorientasi objek.
//Biasanya dalam pemrograman berorientasi objek, ada satu data parent di ouncak yang bisa dianggap sebagai semua implementasi data yang ada di bahasa pembrogrman tsb.
//Di Golang kita dapat menggunakan interface kosong, Interface kosong adalah interface yang tidak memiliki deklarasi method satupun, hal ini membuat secara otomatis semua tipe data akan mejadi implementasinya.
//contoh :
// fmt.Println(a...interface{})
// panic(a interface{})
// recover()interface{}
// dan lainnya
*/

package main

import "fmt"

// func Ups(i interface{}) interface{} { juga bisa
func Ups(i int) interface{} {
	if i == 1 {
		return 1
	}else if i == 2 {
		return true
	}else{
		return "Ups"
	}
}


func main () {
	var data interface{} = Ups(1)
	fmt.Println(data)
}
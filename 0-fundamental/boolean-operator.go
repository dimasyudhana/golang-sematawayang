package main

import "fmt"

func main() {
	var ujian = 91
	var absensi = 91

	var lulusUjian = ujian >= 90
	var lulusAbsensi = absensi >= 90

	fmt.Println(lulusUjian)
	fmt.Println(lulusAbsensi)

	var lulus = lulusAbsensi && lulusUjian
	fmt.Println(lulus)
}

package main

import "fmt"

func main() {
	var name string

	name = "Bro code"
	fmt.Println(name)

	name = "its bro code"
	fmt.Println(name)

	//bisa deklarasi langsung
	var number = 1234
	fmt.Println(number)

	//versi yang lebih simple menggunakan ":"
	angka := 999
	fmt.Println(angka)

	//ketika sudah di deklarasikan, ga harus pakai ":" lagi
	angka = 777
	fmt.Println(angka)

	//versi yang lebih rapi
	var (
		nama_depan    = "Bro"
		nama_belakang = "code"
		angka_awal    = 123
	)

	fmt.Println(nama_depan, nama_belakang, angka_awal)
}

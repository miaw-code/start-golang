package main

import "fmt"

func main() {

	//mengubah tipe data dari variabel
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)

	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16) //terjadi number overflow karena kelebihan kapasitas

	var name = "Bro Code"
	var e uint8 = name[0] //mengambil value dari huruf pertama name, dan bernilai unsign integer, dan mengkonversi ke bit
	var eString = string(e)

	fmt.Println(name) //print bro code
	fmt.Println(e)    //
	fmt.Println(eString)

}

package main

import "fmt"

func main() {
	type nik string //membuat alias variable

	var noKTP nik = "12345678" //alias nik(string) dipakai untuk variable noKTP
	fmt.Println(noKTP)
}

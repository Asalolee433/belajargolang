package main

import "fmt"

func main() {
	//Mendeklarasikan variabel menggunakan var + nama var + tipe data
	var nama string = "Alex"
	var umur int = 19
	//Tapi juga bisa tanpa var dan golang bisa mengetahui tipe data nya
	city := "Jakarta"
	year := 2025
	club := "Manchester United"

	const pi = 3.14

	fmt.Println("Nama:", nama)
	fmt.Println("Umur:", umur)
	fmt.Println("City:", city)
	fmt.Println("Year:", year)
	fmt.Println("Pi:", pi)
	fmt.Println("Club:", club)
}

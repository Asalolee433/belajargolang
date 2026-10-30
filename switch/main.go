package main

import "fmt"

func main() {
	angka := 10
	switch angka {
	case 1:
		fmt.Println("Angka satu")
	case 2:
		fmt.Println("Angka dua")
	case 3:
		fmt.Println("Angka tiga")
	default:
		fmt.Println("Angka tidak dikenal")
	}
}
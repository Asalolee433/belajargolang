package main

import "fmt"

//Fungsi variadic untuk menjumlahkan semua angka 
func JumlahAngka(angka ...int) int {
	total := 0 
	for _, v := range angka {
		total += v // total + v
	}	
	return total
}

func main() {
	fmt.Println(JumlahAngka(1, 2, 3))		// 6
	fmt.Println(JumlahAngka(10, 20, 30, 5))	// 65
	fmt.Println(JumlahAngka())				// 0 
}
package main

import "fmt"

func main() {
	var angka [3]int = [3]int{10, 20, 30}
	fmt.Println(angka)	// [10 20 30]	
	fmt.Println(angka[1])	// 20
	angka[1] = 80 // Mengubah nilai array
	fmt.Println("Ini adalah index ke-1:", angka[1]) // 80
}

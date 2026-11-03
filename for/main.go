package main

import "fmt"

func main(){
	//Loop dari 1 sampai 5
	for i := 1; i <= 5; i++{
		fmt.Println("Angka ke-", 1)
	}

	a := 1
	for a <= 5 {
		fmt.Println("Iterasi ke-", a)
		a++
	} 

	buah := []string{"apel", "buah", "jeruk"}

	for index, item := range buah{
		fmt.Printf("Index %d = %s\n", index, item) 	// Index 0 = apel
													// Index 1 = buah
													// Index 2 = jeruk					
	}
}
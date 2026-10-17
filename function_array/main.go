package main

import "fmt"

func main() {
	number := [5]int{10, 20, 30, 40, 50}

	fmt.Println("Jumlah elemen:", len(number))
	fmt.Println("Index ke 1:", number[1])
	number[1] = 80
	fmt.Println("Index ke 1:", number[1])

	fmt.Println("Ini adalah array :")

	for index , value := range number{
		fmt.Println("Isi index ke :", index," = ",value)
	}

	arr1 := [3]int{1,2,3}
	arr2 := [3]int{4,5,6}

	fmt.Println("arr1 == arr2:", arr1==arr2) //false karena == artinya apakah array sama
	fmt.Println("arr1 != arr2:", arr1 != arr2) //true karna != artinya apakah array berbeda
}

package main

import "fmt"

func main() {
	s := make([]int, 3, 5)
	fmt.Println(s)
	fmt.Println("Len :", len(s))
	fmt.Println("Cap :", cap(s))

	s = append(s, 10, 20)
	fmt.Println("Setelah di append :", s)

	slice2 := make([]int,3)
	slice3 := copy(slice2, s)

	fmt.Println("copied :", slice2)
	fmt.Println("jumlah elemen yang tersalin :", slice3)

	angka := []int{1,2,3,4,5}
	slice4 := angka[1:4]
	fmt.Println("Slice 4:", slice4)
}

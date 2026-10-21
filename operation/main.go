package main

import "fmt"

func main() {
	x := 10
	x = 5
	fmt.Println(x) //5

	x += 2 //x = x+2
	fmt.Println("x +=2 :", x) //7

	x -= 2 //x = x-2
	fmt.Println("x -=2 :", x) //5, karena x yang dicetak adalah hasil x yang di operasikan diatas yaitu 5+2=7 (-2=5)
							  //berlaku untuk seterusnnya(ke bawah)

	x *= 4 //x = x*4
	fmt.Println("x *=4 :", x) //20

	x /= 2 //x = x/2
	fmt.Println("x /=2 :", x) //10

	x %= 2 //x = x%2
	fmt.Println("x %=2 :", x) //0
}

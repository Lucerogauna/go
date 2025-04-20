package main

import "fmt"

const tope = 3

func for1() {
	suma := 0
	num := 0
	for num <= 250 {
		if num%2 == 0 {
			fmt.Println(num)
			suma += num
		}

		num++

	}
	fmt.Println(suma)
}

func downTo() {

	suma := 0
	num := 250
	for num > 0 {
		if num%2 == 0 {
			fmt.Println(num)
			suma += num
		}

		num--

	}
	fmt.Println(suma)

}
func usandoConst() {
	num := 0
	for num < tope {

		fmt.Println("imprimir")
		fmt.Println(num)
		num++
	}
}
func main() {
	fmt.Println("hace usando down to")
	downTo()
	fmt.Println("hace usando for ")
	for1()
	fmt.Println("hace usando const ")
	usandoConst()
}

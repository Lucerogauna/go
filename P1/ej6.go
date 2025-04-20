package main

import "fmt"

func ingresa(x *int, y *int) {
	fmt.Println("leer un numero")
	fmt.Scan(x)
	fmt.Println("leer otro un numero")
	fmt.Scan(y)
	/*
	o puedo hacer
	fmt.Println("leer 2 numeros")
	fmt.Scan(x, y)
	*/
	fmt.Println("Se leyo los numeros")
	fmt.Println(*x, *y)

}

func analiza(x *int, y *int) {
	var z int
	if *x > *y {
		z = *x / *y
	} else {
		z = *y / *x
	}
	fmt.Println("resultado")
	fmt.Println(z)
}
func main() {
	var a, b int
	ingresa(&a, &b)
	analiza(&a, &b)
}

package main

import "fmt"

func ingresa(x *float64, y *float64) {
	fmt.Println("leer un numero")
	fmt.Scan(x)
	fmt.Println("leer otro un numero")
	fmt.Scan(y)
	fmt.Println("Se leyo los numeros")
	fmt.Println(*x, *y)

}

func analiza(x *float64, y *float64) {
	var z float64
	if *x > *y {
		z = *x / *y
	} else {
		z = *y / *x
	}
	fmt.Println("resultado")
	fmt.Println(z)
}
func main() {
	var a, b float64;
	ingresa(&a, &b)
	analiza(&a, &b)
}
//+Inf significa infinito positivo. Es un valor especial que aparece cuando hacés una operación con float64 que no tiene un resultado finito
/*
para ver si una variable es +inf uso 
math.IsInf()
*/
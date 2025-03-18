package main

import "fmt" //paquete fmt
var k, j int
var r bool

func main() {
	fmt.Println("Hello world") //funcion del paquete fmt para imprimir
	var i int                  //declaracion de variables
	i = 6                      //asignacion de variable
	k := 3                     //asignacion directa - short assignment
	k = 38648                  // pisa la asignacion directa y cambia el valor
	r = true
	p1 := "hola"
	p2 := "user"
	fmt.Println(i)
	fmt.Println(k)
	fmt.Println(r)
	fmt.Println(p1 + p2)
}

/*
7. Las temperaturas de los pacientes de un hospital se dividen en 3
grupos: alta (mayor de 37.5), normal (entre 36 y 37.5) y baja
(menor de 36). Se deben leer 10 temperaturas de pacientes e
informar el porcentaje de pacientes de cada grupo. Luego se
debe imprimir el promedio entero entre la temperatura máxima y
la temperatura mínima.
a. ¿Se puede utilizar el case para tipos reales en otros
lenguajes?
b. ¿ Cómo se realizan las conversiones entre reales (punto
flotante) y enteros en otros lenguajes ?
Sub-objetivo: El tipado fuerte, usar casting. Operaciones y
E/S con float. Casting en otros lenguajes.
*/

package main

import "fmt"

func main(){
	var num float64
	var alta, normal, baja, min, max float64
	min = 9999;
	max=-1
	for  i :=0; i<10; i++{
		fmt.Println("leer temperatura")
		fmt.Scan(&num)
	switch {
	case num>37.5 :
		alta ++
	case num < 36:
		baja ++
	default:
		normal++ 
	}
	if  num < min{
		min = num
	}
	 if num > max{
		max = num
	}

}
fmt.Println("promedio de temperatura alta", alta/float64(10))
promedioEntero := int((max + min) / 2)
fmt.Printf("Promedio entero entre la temperatura máxima y mínima: %d\n", promedioEntero)
}
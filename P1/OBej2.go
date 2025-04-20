 //preguntar!! no guarda la ultima palabra si es impat!
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var cadena string
	fmt.Println("Leer una frase: ")
	aux:= [] rune{};
	text := []rune{};
// hola que tal 
	// Creamos un lector para leer la línea completa (incluyendo espacios)
	reader := bufio.NewReader(os.Stdin)
	cadena, _ = reader.ReadString('\n') //subestima el error
	rcadena := [] rune(cadena)
	fmt.Println("La frase ingresada fue:", cadena)
	i:= 0


	for i<len(rcadena){ // recorro toda la cadena
 //for de abajo es un while iterador, no existe while en go
	
		aux= []rune{} // seteo la runa con 0 elementos para que en cada iteracion tenga la cantidad de caracteres de la palabra
		for i<len(rcadena)&& rcadena[i] != ' ' && rcadena[i]!=  '\n'{
			aux= append(aux , rcadena[i])
			i++;
		}
		if len(aux) %2 == 0 { // chequeo si la longitud es par
			text = append(text, aux...)
		}else{
			aux2 := []rune{}; // runa para guardar al reves 
			for j:=len(aux)-1;j>=0; j--{ //o aux-1?
				aux2= append(aux2, aux[j]) //invierto 
			}
			text = append(text, aux2...)
		}
		// Agrega espacio entre palabras (si corresponde)
		 if i < len(rcadena) && rcadena[i] == ' ' {
			text = append(text, ' ')
		}
		i++
	}



	fmt.Println("resultado es", string(text))
}
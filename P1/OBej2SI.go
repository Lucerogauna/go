package main
// VER STRINGS.FIELD 
//split
import (
	"bufio"
	"fmt"
	"strings"
	"os"
)

func invertir (palabra string) string{
	aux := [] rune (palabra)
	aux2 := []rune{}; // runa para guardar al reves 
	for j:=len(aux)-1;j>=0; j--{ 
		aux2= append(aux2, aux[j]) //invierto 
	}
	return string(aux2)
}


func main() {
	var cadena string
	fmt.Println("Leer una frase: ")
	reader := bufio.NewReader(os.Stdin)
	cadena, _ = reader.ReadString('\n') //subestima el error
//cadena:= "hola soy juana y ando en patineta"
	rcadenas := strings.Fields(cadena)
	contador:=0
	for i, palabra:=  range rcadenas {
		contador++;
		if !(contador%2 == 0) {
			rcadenas[i] = invertir (palabra)
		}
		
	}
	nueva:= strings.Join(rcadenas, " ")
	fmt.Println("la nueva cadena es" , nueva)
}


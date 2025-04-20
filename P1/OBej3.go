/*3. Realice un programa que reciba una palabra como argumento y lee de
la entrada una frase. Luego, el programa debe imprimir la frase que
leyó con cada una de las ocurrencias de la palabra con las mayúsculas
y minúsculas invertidas. Por ejemplo, si la frase es:
“Parece peqUEño, pero no es tan pequeÑo el PEQUEÑO”
y la palabra es “PEQUEÑO” entonces el programa imprimirá:
“Parece PEQueÑO, pero no es tan PEQUEñO el pequeño”
Tenga en cuenta que la palabra a buscar puede ser ingresada con
mayúsculas y minúsculas mezcladas.
Este último ejercicio es el que deben entregar.*/
// QUE HAGO CUANDO ME ENCUENTRO CON UN SALTO DE LINEA LO GUARDO?
package main
import ("fmt"
"bufio"
"os"
"unicode")
func main (){
	var word string
	fmt.Println("lee una palabra")
	fmt.Scan (&word)


	fmt.Println("lee una frase")
	reader:= bufio.NewReader(os.Stdin)
	frase, _ = reader.ReadString('\n') //lee la frase por teclado
	
	palabra :=  rune (word)
	runaFrase :=[] rune (frase)
	
	aux :=  [] rune {}
	 i:= 0

	for i<len(frase){ // recorro toda la frase
		aux =  [] rune {} //runa auxiliar que tiene una palabra hasta un espacio
		inicio := i // para reemplazar desde inicio hasta len
		for runaFrase[i]!= '' && runaFrase[i]!= '\n'{
			aux = append (aux, runaFrase[i])
			i++
			} //aca ya tiene la palabra en aux, 
		 //compara las palanras
			if strings.EqualFold(string(aux),string (palabra)){ // o word. si es igual a la que estoy buscando
			for pos :=0;pos<len(aux);pos++{ 
				if unicode.IsUpper(aux[i]) && unicode.IsLower(palabra[i]){
					aux[pos] =unicode.ToLower(aux[i]) // reemplazo 
					}else {
						aux[pos] = unicode.ToUpper(aux[i])
					}
				}
			} //del if
			//reemplazo en la frase original
			for j:= 0; j<len(aux); j++{
				runaFrase[inicio] = aux[j]
			}

	} //del  for principal
	fmt.Println("Frase modificada:")
	fmt.Println(string(runaFrase))
}//del main
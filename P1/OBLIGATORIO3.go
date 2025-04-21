package main
import ("fmt"
"bufio"
"os"
"strings"
"unicode")


func modificarMayusculasYMinusculas(palabra string) string {
	aux := [] rune (palabra)
	for i:= 0; i<len(aux); i++ {
		if unicode.IsLower(aux[i]){
			aux[i]= unicode.ToUpper(aux[i])
		} else{
			aux[i]= unicode.ToLower(aux[i])
		}
	}
	return string(aux)
}
func main (){
	//var word string
	var frase string
	reader:= bufio.NewReader(os.Stdin)
    word := os.Args[1]

    // You can get individual args with normal indexing.
    //arg := os.Args[1]

    fmt.Println(word)
   
	fmt.Println("lee una frase")
	frase, _ = reader.ReadString('\n') //lee la frase por teclado
	
	partes := strings.Fields(frase)
	for i, palabra := range partes{
		if strings.EqualFold(palabra, word){ // o word. si es igual a la que estoy buscando sin importar min y may
			partes[i] = modificarMayusculasYMinusculas(palabra)
		}

	} //del  for principal
	                                                                                                                                
	nueva := strings.Join(partes, " ")
	fmt.Println("Frase modificada:")
	fmt.Println(nueva)
}//del main
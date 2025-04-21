package main
import ("fmt"
"bufio"
"os"
"strings"
"unicode")

func limpiarPalabra(palabra string) (letras string, signos string) {
	runes := []rune(palabra)
	i := len(runes) - 1 //Hola.. i=5 // desde el final de la palabra
	for i >= 0 && !unicode.IsLetter(runes[i]) { //se ejecuta mientras sea un signo
		i--
	}
	return string(runes[:i+1]), string(runes[i+1:])
	//1 devuelve desde el principio hasta la ultima letra inclusive
	//2 desde e caracter hasta el final
}

func pasarMinuscula (palabra string) string{
	aux := [] rune (palabra)
	for i:= 0; i<len(aux); i++ {
		if unicode.IsUpper(aux[i]){
			aux[i]= unicode.ToLower(aux[i])
		}
	}
	return string(aux)
}
func modificarMayusculasYMinusculas(palabra string) string {
	aux := [] rune (palabra)
	for i:= 0; i<len(aux); i++ {
		if unicode.IsLower(aux[i]){
			aux[i]= unicode.ToUpper(aux[i])
		} else if unicode.IsUpper(aux[i]) {
			aux[i]= unicode.ToLower(aux[i])
		}
		//si no cumple ninguna lo guardo como esta sin modificar
	}
	return string(aux)
}
func main (){
	//var word string
	var frase string
	reader:= bufio.NewReader(os.Stdin)
    word := os.Args[1] //guardo la palabra como arg 
	fmt.Println(word)
    // You can get individual args with normal indexing.

   
	fmt.Println("lee una frase")
	frase, _ = reader.ReadString('\n') //lee la frase por teclado

	partes := strings.Fields(frase)
	//wordEnMin := pasarMinuscula(word) //  caso limite por ej: quiero ingresar una palabra en mayuscula 

	for i, palabra := range partes{
		//divido la palabra en 2 strings, uno con la palabra limpia y el otro solo signos

		palabraLimpia, signos:= limpiarPalabra(palabra)

		//OTRA OPCION if  pasarMinuscula(palabraLimpia) == wordEnMin{ // a la palabra la paso en minuscula para comparar
		//con equalFold se ve mas limpio y no hace falta declarar variables extras como wordEnMin
		
		if strings.EqualFold(palabraLimpia, word){ // si la palabra limpia es igual a la palabra pasada por arg (sin importar las min y mayus)
			palabraLimpia = modificarMayusculasYMinusculas(palabraLimpia) 
			partes[i] = palabraLimpia + signos // guardo la palabra de esa posicion ya modificada
		}
	} //del  for principal
	                                                                                                                                
	nueva := strings.Join(partes, " ")
	fmt.Println("Frase modificada:")
	fmt.Println(nueva)
}//del main
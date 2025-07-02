package main
import ("fmt"
"bufio"
"os"
"strings"
"unicode")

func limpiarPalabra(palabra string) (letras string, signos string, atras bool) {
	runes := []rune(palabra)
	if !unicode.IsLetter(runes[len(runes)-1]){ // si tiene un caracter especial la ultima pos
		i := len(runes) - 1 //Hola.. i=5 // desde el final de la palabra
		for i >= 0 && !unicode.IsLetter(runes[i]) { //se ejecuta mientras sea un signo
			i--
		}
		return string(runes[:i+1]), string(runes[i+1:]), true
	//1 devuelve desde el principio hasta la ultima letra inclusive
	//2 desde e caracter hasta el final
	//3 un true porque estaban atras
	}
	 if !unicode.IsLetter(runes[0]){ // me fijo si tiene puntos/signos adelante
		i:= 0
		for i <len(runes) && !unicode.IsLetter(runes[i]){
			i++
		}
		return string(runes[i:]), string(runes[:i]), false
	}
	//1 devuelve desde el principio hasta la ultima letra inclusive
	//2 desde e caracter hasta el final
	//3 dalse porque los signos estan adelante
}
	// si no tiene signos ni adelante ni atrás, la devuelvo tal cual
	return palabra, "", true

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
    word := os.Args[1] //guardo la palabra que viene como argumento por consola
	fmt.Println(word)

	
	fmt.Println("lee una frase")
	frase, _ = reader.ReadString('\n') //lee la frase por teclado

	partes := strings.Fields(frase)

	for i, palabra := range partes{
		//divido la palabra en 2 strings, uno con la palabra limpia y el otro solo signos

		palabraLimpia, signos, estaAlFinal:= limpiarPalabra(palabra)

		 //equals compara la palabra sin importar las min y mayus
		if strings.EqualFold(palabraLimpia, word){ // si la palabra limpia es igual a la palabra pasada por arg (sin importar las min y mayus)
			palabraLimpia = modificarMayusculasYMinusculas(palabraLimpia) 
			//vuelvo a armar la palabra con los signos 
			if (estaAlFinal == true){
				partes[i] = palabraLimpia + signos // guardo la palabra de esa posicion ya modificada
			}else{
				partes[i] = signos + palabraLimpia }	
			}
	} //del  for principal
	                                                                                                                                
	nueva := strings.Join(partes, " ")
	fmt.Println("Frase modificada:")
	fmt.Println(nueva)
}//del main

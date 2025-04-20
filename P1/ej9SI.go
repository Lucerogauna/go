//indexOfString
package main
import ("fmt" 
		"strings"
		"unicode"
)
func cambiarPorMartes(palabra, reemplazo string) string {
    // Convierte el string 'reemplazo' respetando las mayúsculas de 'original'
    resultado := []rune(reemplazo)
    for i, palab := range palabra {
        if i < len(resultado) {
            if unicode.IsUpper(palab) {
                resultado[i] = unicode.ToUpper(resultado[i])
            } else {
                resultado[i] = unicode.ToLower(resultado[i])
            }
        }
    }
    return string(resultado)
}
func main (){
	frase := "hola soy jueves Jueves JuEvES JUEVES" // vector de caracteres
	fmt.Println(frase)
	palabras := strings.Fields(frase) //te armaria un slice por palabra

    //fmt.Println("Vector de palabras:")
	for i, palabra := range palabras {
		if strings.EqualFold(palabra, "jueves") {
			//fmt.Println("Se leyó 'jueves' en la posición", i)
			palabras[i]= cambiarPorMartes(palabra, "martes")

		}
	}
	nuevaFrase := strings.Join(palabras, " ") //unir los elementos de un slice de strings en un solo string, usando un separador que vos elijas.


	fmt.Println("Modificada:", nuevaFrase)

}	
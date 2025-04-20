package main
import ("fmt" 
		"strings"
		"unicode"
)
func cambiar( orig, remp string) string {
resultado := []rune(remp)
for i, palab := range orig {
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
func main() {
    frase := "los jueves voy en martes Miercoles"
    fmt.Println("Frase original:", frase)

    palabras := strings.Fields(frase)

    for i, palabra := range palabras {
        if strings.EqualFold(palabra, "jueves") {
            palabras[i] = cambiar(palabra, "martes")
        } else if strings.EqualFold(palabra, "miercoles") {
            palabras[i] = cambiar(palabra, "automovil")
        }
    }

    nuevaFrase := strings.Join(palabras, " ")
    fmt.Println("Frase modificada:", nuevaFrase)
}
package main
import ("fmt"
 "strings"
 "bufio"
 "os")


func main(){
	fmt.Println(strings.ToUpper("Gopher")) // devuelve todas las letras a mayusculas
	fmt.Println(strings.ToLower("GOPHER")) // Devuelve todas las letras mayusculas a minusculas
	fmt.Println(strings.ReplaceAll("oink oink oink", "oink", "moo")) // string, viejo, nuevo 

	fmt.Println(strings.Count("cheese", "e")) //  cuenta el número de instancias no superpuestas de substr en s
	fmt.Println(strings.Contains("seafood", "foo")) //si substr está dentro de s.
	fmt.Println(strings.EqualFold("Go", "go")) //insensibilidad a las min y mayus, devuelve si son iguales
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Escribí una frase:")
	
	
	
	///Preguntar lo de error  
	frase, _ := reader.ReadString('\n') //Lee todo lo que escriba el usuario hasta que presione Enter ('\n').
	//frase = strings.TrimSpace(frase) // preguntar si hay quehacerlo
	fmt.Println(frase)
}

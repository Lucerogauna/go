package main
import "fmt"
//map es = que tabla de hash
func main(){
type Map[K comparable, V any] map[K]V
var mapa Map[string, int]= make (Map[string, int])// crea el mapa con clave y valor
mapa ["Lucero"]=22
mapa["Nahuel"]= 21
mapa["Pedro"]= 24
mapa["Lucio"]=21
fmt.Println("mapa")
for nombre, edad:= range mapa{
		fmt.Printf("%s: %d años\n", nombre, edad)
	
}
fmt.Println(mapa)

mapa2:= Map[string, []string]{
"Fod": {"Lucero", "Lucia", "Lucio", "Luciano"},
"Ayed": {"Marga"},
"Go": {"Gonzalo", "Carmen"},
}
fmt.Println(mapa2)

}


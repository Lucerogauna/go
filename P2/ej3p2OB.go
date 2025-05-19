package main

type Elemento struct {
	Nro int
	CanttOcurrencias int
}

type OptimumSlice struct {
	Info []Elemento
}

//inciso i
func New(s slice) OptimumSlice{
	if len(slice)==0{
		return OptimumSlice{} // si no tiene elementos retorta os sin elementos
	}
	var resultado OptimumSlice
	nroAct:= slice[0]
	cont:= 1;

	cont (i:= 1; i<len(slice); i++){
		if slice[i]==nroAct{
			cont++
		}else{
		//leyo otro numero<> a nroAct
		//Crea un Elemento con ese número y su cantidad de repeticiones, y lo mete al final de la lista comprimida:
			resultado.Info= append (resultado.Datos, Elemento{nroAct,cont})
			//actualizo con el nuevo numero
			nroAct= slice[i]
			cont=1
		}
	}
//agrego el ultimo
resultado.Info= append(resultado.Info, Elemento{nroAct,cont})
return resultado
}




//inciso ii
func IsEmpty(o OptimumSlice) bool {
	return len(o.Info) == 0
}



//inciso iii
func Len(o OptimumSlice)int{


}


//inciso iv
func FrontElement(o OptimumSlice) int {
	if IsEmpty(o){
		return -1 // o algo que indique vacío
	}
	return o.Info[0].Nro
}

//inciso v
func LastElement(o OptimumSlice) int {
	if IsEmpty(o){
		return -1 // o algo que indique vacío
	}

	return o.Info[len(o)-1].Nro
}



//inciso vi
func Insert(o OptimumSlice, element int, position int) int


func main(){

}

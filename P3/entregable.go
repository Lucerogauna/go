package main
import ("fmt"
"sync"
)

type Tarea struct {
    Numero   int
    Prioridad int
}
func generarTareas(cantidad int) []Tarea {
	rand.Seed(time.Now().UnixNano())
	tareas := make([]Tarea, 0, cantidad)
	for i := 0; i < cantidad; i++ {
		num := rand.Intn(900) + 100 // Entre 100 y 999
		prioridad := rand.Intn(4)   // Entre 0 y 3
		tareas = append(tareas, Tarea{num, prioridad})
	}
	return tareas
}

func hacerTrabajo(prioridad int, tareas <-chan Tarea, wg *sync.WaitGroup, mutex *sync.Mutex, acumulador *int, archivo *os.File){
	
}
func main(){

	const cantidadTareas = 10
	tareasGeneradas := generarTareas(cantidadTareas)
	// Canales por prioridad
	canales := []chan Tarea{
		make(chan Tarea, cantidadTareas),
		make(chan Tarea, cantidadTareas),
		make(chan Tarea, cantidadTareas),
		make(chan Tarea, cantidadTareas),
	}

	// PUNTO F
	// Archivos para prioridades 0 y 1
	archivo0, _ := os.Create("prioridad0.txt")
	defer archivo0.Close()
	archivo1, _ := os.Create("prioridad1.txt")
	
	var wg sync.WaitGroup //crea a un waitgroup que espera a que los 4 terminen antes de seguir
	wg.Add(4)
	for i := 0; i < 4; i++ {
    var archivo *os.File
    if i == 0 {
        archivo = archivo0
    } else if i == 1 {
        archivo = archivo1
    }
    go hacerTrabajo(i, canales[i], &wg, &mutex, &acumulador, archivo)
}

	defer archivo1.Close()
	//acumulador para prioridad 3
	var acumulador int
	var mutex sync.Mutex
	//distribuye tarea por prioridad
	for _, t := range tareasGeneradas{
		canales[t.Prioridad] <-t
	}
	//cierra todos los canales
	for _, ch := range canales {
		close(ch)
	}

	wg.Wait()
	fmt.Println("Todas las tareas fueron procesadas.")
	}
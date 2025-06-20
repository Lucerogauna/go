/*Queremos que tres goroutines impriman un mensajito cada una,
 y el programa espere a que terminen todas antes de cerrarse.
*/
package main

import (
	"fmt"
	"sync"
)

func main() {

/*wg.Done() se comunica con wg.Wait()
Wait() espera que se hayan llamado
 todas las Done() necesarias para continuar.*/

	var wg sync.WaitGroup // Creamos un WaitGroup

	// Lanzamos 3 tareas (goroutines)
	for i := 1; i <= 5; i++ {
		wg.Add(1) // Avisamos que se agrega una tarea pendiente

		go func(n int) {
			defer wg.Done() // se resta 1 al contador ✅ Siempre se va a ejecutar cuando termine esta goroutine

			fmt.Printf("Hola, soy la goroutine #%d\n", n)
		}(i) // Le pasamos el número de goroutine como parámetro
	}

	wg.Wait() // Esperamos a que todas las goroutines terminen
	fmt.Println("¡Todas las tareas terminaron!")
}

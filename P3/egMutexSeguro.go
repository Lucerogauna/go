package main

import (
	"fmt"
	"sync"
	"time"
)

// Estructura con contador protegido por RWMutex
type ContadorSeguro struct {
	valor int
	mu    sync.RWMutex
}

// Método que suma al contador (bloquea escritura)
func (c *ContadorSeguro) Sumar() {
	c.mu.Lock()         // 🔒 Candado de escritura (solo uno puede escribir)
	defer c.mu.Unlock() // 🔓 Lo libera cuando termina
/*defer en Go registra una acción para que se ejecute automáticamente al final de la función (cuando termina, ya sea por return, error, etc.).

No se ejecuta ahí mismo, sino más tarde, cuando la función se está por ir.*/
	c.valor++
}

// Método que lee el valor actual (permite varias lecturas)
func (c *ContadorSeguro) Leer() int {
	c.mu.RLock()         // 🔒 Candado de lectura (pueden entrar varios a leer)
	defer c.mu.RUnlock() // 🔓 Lo libera cuando termina

	return c.valor
}

func main() {
	var wg sync.WaitGroup
	/*Es una herramienta que sirve para esperar a que un grupo de goroutines termine antes de seguir con el programa.*/
	contador := ContadorSeguro{}

	// Lanzamos 5 goroutines que suman al contador
	for i := 0; i < 5; i++ {
		wg.Add(1) // Se le dice: “voy a lanzar 1 tarea” (puede ser más)
		go func() {

			for j := 0; j < 1000; j++ {
				//codigo paralelo
				contador.Sumar()
			}
			defer wg.Done() // o uso wg.Done solo
		}()
	}

	// Goroutine que lee el valor cada medio segundo
	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(500 * time.Millisecond)
			fmt.Println("Valor actual del contador:", contador.Leer())
		}
	}()

	wg.Wait() // Esperamos a que terminen todas las sumas

	// Mostramos el valor final
	fmt.Println("Valor final del contador:", contador.Leer())
}

/*Get-Content temperaturas.txt | go run ej2p2.go
Esto envía el contenido del archivo temperaturas.txt como entrada estándar al programa de Go.*/
package main
import "fmt"

func main() {
	const N = 4
	var temperaturas[N] float64         // suma de temperaturas por categoría
	var contadores[N] int               // cantidad de temperaturas por categoría
	var max, min float64
	min = 999.9
	max = 0.0
	var temperatura float64

		// Leer hasta EOF
	for {
		_, err := fmt.Scan(&temperatura)
		if err != nil {
			break // fin de datos
		}

		if temperatura > 37.5 && temperatura < 50 {
			temperaturas[0] += temperatura // temp altas
			contadores[0]++
		} else if temperatura >= 36 {
			temperaturas[1] += temperatura // temp normales
			contadores[1]++
		} else if temperatura < 36 && temperatura > 20 {
			temperaturas[2] += temperatura // temp bajas
			contadores[2]++
		} else {
			temperaturas[3] += temperatura // temp no válidas
			contadores[3]++
		}

		if temperatura > max {
			max = temperatura
		}
		if temperatura < min {
			min = temperatura
		}
	}

	// Mostrar promedios por categoría (si hubo datos)
	categorias := [4]string{"alta", "normal", "baja", "no válida"}
	for i := 0; i < N; i++ {
		if contadores[i] > 0 {
			fmt.Printf("Promedio temperatura %s: %.2f\n", categorias[i], temperaturas[i]/float64(contadores[i]))
		} else {
			fmt.Printf("No se ingresaron temperaturas %s\n", categorias[i])
		}
	}

	// Promedio entero entre max y min
	promedio := int((max + min) / 2)
	fmt.Printf("Promedio entero entre la temperatura máxima y mínima: %d\n", promedio)
}

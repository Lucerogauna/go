package main
improt "fmt"

func main(){
	const N=10
	var temperaturas[N] float64
	var alta,normal,baja int
	var max, min float64
	min = 999
	max=0
// Leer las temperaturas
for i := 0; i < N; i++ {
	fmt.Scan(&temperaturas[i])
	if temperaturas[i] > 37.5 {
		alta++
	} else if temp >= 36 {
		normal++
	} else {
		baja++
	}
	if temperaturas[i] > max {
		max = temperaturas[i]
	}
	if temperaturas[i] < min {
		min = temperaturas[i]
		}

}

// Calcular porcentajes
fmt.Printf("Porcentaje con temperatura alta: %.2f%%\n", float64(alta)*100/float64(N))
fmt.Printf("Porcentaje con temperatura normal: %.2f%%\n", float64(normal)*100/float64(N))
fmt.Printf("Porcentaje con temperatura baja: %.2f%%\n", float64(baja)*100/float64(N))

// Promedio entero entre max y min
promedio := int((max + min) / 2)
fmt.Printf("Promedio entero entre la temperatura máxima y mínima: %d\n", promedio)
}


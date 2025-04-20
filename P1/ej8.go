/*8. Realizar un programa que lea el punto cardinal (como caracter o
	string) del cual viene el viento (‘N’, ‘S’, ‘E’, ‘O’) y envíe a la salida
	estándar hacia cuál se dirigiría.
	Sub-objetivo: Uso de case con la opción por default. E/S
	caracteres o strings.
	a. ¿Cómo se escribe el default en el case de otros lenguajes?*/

	package main
	import "fmt"

	func main (){
		 var c string
		fmt.Println("leer")
		fmt.Scan(&c)
		switch c {
		case "N": fmt.Println("viento direccion Norte")
		case "S": fmt. Println("viento direccion Sur")
		
		case "E": fmt. Println("viento direccion Este")
		case "O": fmt. Println("viento direccion Oeste")
		default:
			fmt.Println("error");			
		}

	}
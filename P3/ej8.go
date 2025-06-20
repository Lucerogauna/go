package main

import ("sync"
"fmt")

// a) Estructura Contacto
type Contact struct {
	Nombre           string
	Apellido         string
	CorreoElectronico string
	Telefono         string
}

// b) Estructura Agenda con protección para concurrencia
type Agenda struct {
	contactos map[string]Contact /*Internamente, Go lo maneja como una tabla hash, así que la "posición" depende de cómo la clave (en este caso, el string del correo) se convierte a una ubicación interna.*/
	mutex       sync.RWMutex // RWMutex para lectura y escritura segura
/*sync.RWMutex es un candado (mutex = mutual exclusion) que permite controlar el acceso concurrente a datos compartidos, como el map de contactos.*/
}

func (a *Agenda) AgregarContacto(contacto Contact){
	a.mutex.Lock()
	defer a.mutex.Unlock()
	a.contactos[contacto.correoElectronico]= contacto
	fmt.Printf("✔️ Agregado: %s\n", contacto.CorreoElectronico)
	
}
func (a *Agenda)EliminarContacto(correo string){
	a.mutex.Lock()
	defer a.mutex.Unlock() // lo ejecuta al final acordarse
	delate(a.contactos, correo)
	/*No es obligatorio buscar antes de eliminar porque:
	delete es seguro: no genera error si la clave no está.*/
	fmt.Printf("❌ Eliminado: %s\n", correo)
}


/*📌 El lenguaje no garantiza seguridad de concurrencia en map incluso para lecturas.

❌ Si una goroutine está escribiendo al mismo tiempo que otra está leyendo, podés tener errores como:
fatal error: concurrent map read and map write*/
func (a *Agenda) BuscarContacto(correo string) (Contact, bool){
	a.mutex.Lock()
	defer a.mutex.Unlock() // lo ejecuta al final acordarse

	contacto, existe := a.contactos[correo]
	if existe {
		fmt.Printf("✅ Se encontró el contacto: %s\n", correo)
	} else {
		fmt.Printf("❌ No se encontró el contacto: %s\n", correo)
	}
	return contacto, existe
}
/* función main() que cree una agenda, inicie varias
goroutines para agregar, eliminar y buscar contactos de manera
simultánea, y luego imprima el contenido de la agenda después
de un tiempo para verificar que las operaciones se hayan
realizado correctamente*/
func main(){
	agenda := Agenda{
	contactos: make(map[string]Contact),
	}

	var wg sync.WaitGroup

	// Lista de contactos para agregar
	contactos := []Contact{
		{"Ana", "García", "ana@gmail.com", "1111"},
		{"Luis", "Martínez", "luis@gmail.com", "2222"},
		{"María", "López", "maria@gmail.com", "3333"},
		{"Juan", "Pérez", "juan@gmail.com", "4444"},
		{"Elena", "Ríos", "elena@gmail.com", "5555"},
	}
	for _, c := range contactos { //for each usando c como cada contacto
		wg.Add(1)//wg.Add(n) → Le decís cuántas goroutines vas a esperar.
		
		/*	Cuando querés hacer más cosas dentro de la goroutine
		 (ej: imprimir, controlar wg, etc.)*/
		go func(c Contact) {
			defer wg.Done() //Llamás esto cuando una goroutine termina.
			agenda.AgregarContacto(c)
			fmt.Println("✔️ Agregado:", c.CorreoElectronico)
		}(c)
		
	/*NOTA Para que se ejecute, tenés que llamar a la función, 
	y la forma es poner (c) justo después de la definición:
	Go func(c Contact) {
    	// código
	}(c)  // <-- esto llama la función pasando el valor c
	// ¿Para qué sirve pasar c como parámetro?
	En un bucle, la variable c cambia en cada iteración. Si no la pasás como parámetro, todas las goroutines podrían terminar usando el mismo valor (el último de la iteración), porque la función anónima "captura" la variable del entorno y no su valor.

	Pasar el valor como parámetro crea una copia local independiente para cada goroutine.*/
	}// Buscar contactos concurrentemente
	emailsBuscar := []string{
		"ana@gmail.com",
		"juan@gmail.com",
		"noexiste@gmail.com",
	}
	for _, email := range emailsBuscar {
		wg.Add(1)
		go func(correo string) {
			defer wg.Done()
			if c, ok := agenda.BuscarContacto(correo); ok {
				fmt.Printf("🔍 Encontrado: %s %s\n", c.Nombre, c.Apellido)
			} else {
				fmt.Printf("🔍 No encontrado: %s\n", correo)
			}
		}(email)
	}

	// Eliminar contactos concurrentemente
	emailsEliminar := []string{
		"ana@gmail.com",
		"elena@gmail.com",
	}
	for _, email := range emailsEliminar {
		wg.Add(1)
		go func(correo string) {
			defer wg.Done()
			agenda.EliminarContacto(correo)
			fmt.Println("❌ Eliminado:", correo)
		}(email)
	}

	// Esperar a que todas las operaciones terminen
	wg.Wait()

	// Esperar un poco más por seguridad (no siempre necesario)

	time.Sleep(500 * time.Millisecond)

	// Imprimir contactos restantes
	agenda.ImprimirTodo()

}
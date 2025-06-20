package main

import "fmt"

// Interfaz de lista enlazada, define las operaciones básicas
type LinkedLister interface {
	add(element int, position int)
	get(position int) int
	delete(position int)
	getSize() int
}

// Nodo individual de la lista enlazada
type Lista struct {
	valor     int
	siguiente *Lista
}

// Estructura de la lista enlazada propiamente dicha
type LinkedList struct {
	primerNodo *Lista // apunta al primer nodo de la lista
	size       int    // cantidad de elementos
}

// Agrega un nuevo elemento en la posición dada
func (l *LinkedList) add(elemento int, position int) {
	//*LinkedList significa que el método recibe un puntero a una LinkedList. Así, los cambios que hagas dentro del método modifican la estructura original, no una copia.
	if position < 0 || position > l.size {
		fmt.Println("Posición inválida")
		return
	}

	newNode := &Lista{valor: elemento}

	// Si se quiere insertar en la primera posición
	if position == 0 {
		newNode.siguiente = l.primerNodo
		l.primerNodo = newNode
		l.size++
		return
	}

	// Buscar el nodo anterior a la posición deseada
	current := l.primerNodo
	for i := 0; i < position-1; i++ {
		current = current.siguiente
	}

	// Insertar el nuevo nodo
	newNode.siguiente = current.siguiente
	current.siguiente = newNode
	l.size++
}

// Devuelve el valor del nodo en la posición indicada
func (l *LinkedList) get(position int) int {
	if position < 0 || position >= l.size {
		fmt.Println("Posición inválida")
		return -1
	}

	current := l.primerNodo
	for i := 0; i < position; i++ {
		current = current.siguiente
	}
	return current.valor
}

// Elimina el nodo en la posición indicada
func (l *LinkedList) delete(position int) {
	if position < 0 || position >= l.size {
		fmt.Println("Posición inválida")
		return
	}

	// Si es el primer nodo
	if position == 0 {
		l.primerNodo = l.primerNodo.siguiente
		l.size--
		return
	}

	// Buscar el nodo anterior al que se quiere eliminar
	current := l.primerNodo
	for i := 0; i < position-1; i++ {
		current = current.siguiente
	}

	// Saltarse el nodo a eliminar
	current.siguiente = current.siguiente.siguiente
	l.size--
}

// Devuelve el tamaño actual de la lista
func (l *LinkedList) getSize() int {
	return l.size
}

// Función principal de prueba
func main() {
	lista := LinkedList{}

	lista.add(10, 0) // Lista: [10]
	lista.add(20, 1) // Lista: [10, 20]
	lista.add(15, 1) // Lista: [10, 15, 20]

	fmt.Println("Tamaño:", lista.getSize()) // 3 elementos

	fmt.Println("Elemento en posición 0:", lista.get(0)) // 10
	fmt.Println("Elemento en posición 1:", lista.get(1)) // 15
	fmt.Println("Elemento en posición 2:", lista.get(2)) // 20

	lista.delete(1) // Elimina el 15 → Lista: [10, 20]
	fmt.Println("Después de borrar en posición 1:")
	fmt.Println("Elemento en posición 1:", lista.get(1)) // 20
	fmt.Println("Tamaño:", lista.getSize())              // 2
}

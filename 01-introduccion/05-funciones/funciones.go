package main

import "fmt"

func saludar(nombre string) {
	fmt.Printf("Hola %s, saludos desde la funcion saludar\n", nombre)
}

func sumar(a, b int) int {
	return a + b
}

func main() {
	//saludar("Nelson")
	fmt.Printf("%d", sumar(5, 8))
}

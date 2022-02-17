package main

import "fmt"

func modificarValor(cadena *string) {
	fmt.Printf("%p\n", cadena)
	*cadena = "Hola desde la funcion"
}

func main() {
	cadena := "Hola mundo de GO"
	fmt.Printf("%p\n", &cadena)
	fmt.Println("Antes de la funcion:", cadena)
	modificarValor(&cadena)
	fmt.Println("Despues de la funcion:", cadena)
}

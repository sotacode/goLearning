package main

import (
	"fmt"
	"strings"
)

//clousures

func repeat(n int) func(cadena string) string {
	return func(cadena string) string {
		return strings.Repeat(cadena, n)
	}
}

func main() {

	repeat3 := repeat(3)

	fmt.Println(repeat3("hola "))
	fmt.Println(repeat3("Sota "))
	//Funcion anonima se pueden asignar a variables
	/* func() {
		fmt.Println("Hola desde func anonima")
	}() */

	/* myfunc := func(nombre string) string {
		return fmt.Sprintf("Hola %s te saludo desde la func anonima", nombre)
	}

	mensaje := myfunc("Sota")
	fmt.Println(mensaje) */
}

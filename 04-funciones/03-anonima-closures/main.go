package main

import "fmt"

func main() {

	//Funcion anonima se pueden asignar a variables
	/* func() {
		fmt.Println("Hola desde func anonima")
	}() */

	myfunc := func(nombre string) string {
		return fmt.Sprintf("Hola %s te saludo desde la func anonima", nombre)
	}

	mensaje := myfunc("Sota")
	fmt.Println(mensaje)
}

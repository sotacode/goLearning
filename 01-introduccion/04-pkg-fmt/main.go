package main

import "fmt"

func main() {
	hola := "Hola"
	mundo := "mundo"

	fmt.Println(hola)
	fmt.Println(mundo)

	nombre := "Sota"
	edad := 24

	fmt.Printf("Hola %s tu edad es %d\n", nombre, edad)

	mensaje := fmt.Sprintf("Hola %s tu edad es %d", nombre, edad)

	fmt.Print(mensaje)
	fmt.Printf("Tipo de dato de mensaje: %T", mensaje)

	fmt.Printf("Ingrese un nombre:")
	fmt.Scanln(&nombre)
	fmt.Printf("\nEl otro nombre es: %s", nombre)
}

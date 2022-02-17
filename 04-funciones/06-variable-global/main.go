package main

import "fmt"

var mensaje string

func funcion1() {
	mensaje = "Hola desde func 1"
	fmt.Println(mensaje)
}

func funcion2() {
	mensaje = "Hola desde func 2"
	fmt.Println(mensaje)
}

func main() {

	mensaje = "Hola desde main"
	fmt.Println(mensaje)

	defer funcion1()
	funcion2()

	fmt.Println(mensaje)

}

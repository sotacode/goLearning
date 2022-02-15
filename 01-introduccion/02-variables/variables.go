package main

import (
	"fmt"
	"reflect"
)

func main() {
	var nombre1 string
	var apellido1 string

	nombre1 = "Nelson"
	apellido1 = "Rivera"

	edad := 24.5

	var ciudad = "Angol"

	var a int
	var b float64
	var c string
	var d bool

	const pi = 3.143234

	fmt.Println("Mi nombre es: ", nombre1, apellido1, " y tengo la edad de ", edad, ", vivo en ", ciudad)

	fmt.Println("Los valores iniciales son: ", a, b, c, d)

	fmt.Println(reflect.TypeOf(pi))
}

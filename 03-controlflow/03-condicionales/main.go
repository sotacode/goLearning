package main

import "fmt"

func main() {

	var consumo, descuento float64
	var datosDescuentos string

	//Entrada de datos
	fmt.Print("Ingrese consumo: ")
	fmt.Scanln(&consumo)

	if consumo >= 0 {
		datosDescuentos = "10%"
		descuento = consumo * 0.1
		fmt.Printf("El precio original es %f, pero con un descuento del %s queda en: %f", consumo, datosDescuentos, consumo-descuento)
	} else {
		fmt.Println("Error")
	}
}

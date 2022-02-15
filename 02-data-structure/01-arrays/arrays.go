package main

import "fmt"

func main() {

	//igual a linea 22
	var numeros [4]int

	numeros[0] = 10
	numeros[1] = 20
	numeros[2] = 30
	numeros[3] = 60

	fmt.Println(numeros)

	//igual a linea 22
	nombres := [4]string{"Victoria", "Isidora", "Nelson", "Sofia"}

	fmt.Println(nombres)

	//Una vez se calcula la cantidad de elementos, no se pueden agregar mas
	colores := [...]string{
		"Rojo",
		"Verde",
		"Blanco",
		"Morado",
	}

	fmt.Println(colores, len(colores))

	//indices definidos
	monedas := [...]string{
		0: "dolar",
		2: "soles",
		3: "pesos",
		8: "euros",
	}

	fmt.Println(monedas, len(monedas))

	//subarray
	subMoneda := monedas[3:9]
	fmt.Println(subMoneda)
}

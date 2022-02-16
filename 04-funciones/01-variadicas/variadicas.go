package main

import "fmt"

//no sabemos cuantos valores vamos a recibir pero sabemos que son enteros
func sumar(nombre string, numeros ...int) (string, int) {

	mensaje := fmt.Sprintf("La suma de %s es:", nombre)
	result := 0
	for _, valor := range numeros {
		result += valor
	}

	return mensaje, result

}

func main() {
	m, result := sumar("Sota", 4, 5, 8, 2)
	fmt.Println(m, result)
}

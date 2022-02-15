package main

import "fmt"

func main() {
	//slicen
	numeros := []int{1, 2, 3}

	fmt.Println(numeros, len(numeros))

	numeros = append(numeros, 4)

	fmt.Println(numeros, len(numeros))

	subSlicen := numeros[:3]
	numeros[2] = 200

	fmt.Println(subSlicen)
	fmt.Println(numeros)

	//Punteros

	//Longitud

	//Capacidad

	meses := []string{"Enero", "Febrero", "Marzo", "Abril", "Mayo"}

	fmt.Printf("Len: %v - Cap: %v - %p ", len(meses), cap(meses), meses)
	meses = append(meses, "Junio")
	fmt.Printf("Len: %v - Cap: %v - %p ", len(meses), cap(meses), meses)

}

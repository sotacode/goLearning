package main

import "fmt"

func main() {
	numeros := make([]int, 3, 3)

	fmt.Println(numeros, len(numeros), cap(numeros))

	numeros[0] = 100

	numeros = append(numeros, 400)

	fmt.Println(numeros, len(numeros), cap(numeros))
}

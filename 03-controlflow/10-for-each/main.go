package main

import "fmt"

func main() {
	nombres := [...]string{"Alex", "Sota", "Isi", "Manuel"}

	for indice, elemento := range nombres {
		fmt.Println(indice, elemento)
	}
}

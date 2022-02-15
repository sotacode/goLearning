package main

import "fmt"

func main() {

	dias := make(map[int]string)

	fmt.Println(dias)

	//agregar datos

	dias[17] = "Lunes"
	dias[18] = "Martes"
	dias[19] = "Miercoles"
	dias[20] = "Jueves"
	dias[21] = "Viernes"
	dias[22] = "Sabado"

	fmt.Println(dias)

	//modify
	dias[22] = "SABADO"
	fmt.Println(dias)

	//DELETE
	delete(dias, 19)
	fmt.Println(dias)

	//Nueva mapa

	estudiantes := make(map[string][]int)

	estudiantes["Sota"] = []int{20, 21, 22}
	estudiantes["Isi"] = []int{12, 13, 14, 15, 16}

	fmt.Println(estudiantes)
	fmt.Println(estudiantes["Sota"][0])

}

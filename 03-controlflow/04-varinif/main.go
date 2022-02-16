package main

import "fmt"

func main() {

	//variables solo se pueden usar en el scope del if
	if nombre, edad := "Sota", 24; nombre == "Sota" {
		fmt.Println("Sota tiene la edad", edad)
	}

	users := make(map[string]string)
	users["Sota"] = "holarina1@gmail.com"
	users["Isi"] = "isi@gmail.com"

	//podemos ver si un valor de una mapa existe o no de la sgte manera
	//correo, ok := users["Sota"]
	//fmt.Println(correo, ok)

	if correo, ok := users["alex"]; ok {
		fmt.Println("El dato existe y su correo es:", correo)
	} else {
		fmt.Println("No existe la clave")
	}

}

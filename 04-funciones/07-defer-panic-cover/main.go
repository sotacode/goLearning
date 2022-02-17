package main

import (
	"fmt"
	"os"
)

func main() {

	//funcion recover
	defer func() {
		if error := recover(); error != nil {
			fmt.Println("Ups al parecer el programa no finalizo de forma correcta.")
		}
	}()

	if file, error := os.Open("jaja.txt"); error == nil {
		defer func() {
			fmt.Println("Cerrando el archivo")
			file.Close()
		}()
		contenido := make([]byte, 254)
		long, _ := file.Read(contenido)
		contenidoArchivo := string(contenido[:long])
		fmt.Println(contenidoArchivo)
	} else {
		panic("Error al leer el archivo")
	}
}

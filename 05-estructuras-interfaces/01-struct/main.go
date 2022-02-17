package main

import "fmt"

type Persona struct {
	nombre string
	edad   int
}

func (p *Persona) imprimir() {
	fmt.Println(p.nombre, p.edad)
}
func (p *Persona) nuevoNombre(nuevoNombre string) {
	p.nombre = nuevoNombre
}

func main() {
	p1 := Persona{"Nelsota", 24}
	p1.imprimir()
	p1.nuevoNombre("Sota")
	p1.imprimir()
}

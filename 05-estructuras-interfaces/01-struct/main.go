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

//Herencia
type Empleado struct {
	Persona
	sueldo float64
}

func main() {
	p1 := Persona{"Nelsota", 24}
	//p1.imprimir()
	//p1.nuevoNombre("Sota")
	p1.imprimir()

	e1 := Empleado{
		sueldo: 500.0,
	}
	e1.nombre = "Pedro"
	e1.edad = 30
	e1.imprimir()
	fmt.Println(e1)

}

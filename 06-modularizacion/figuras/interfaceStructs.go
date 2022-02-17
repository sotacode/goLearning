package figuras

import "fmt"

type Figuras interface {
	area() float64
	perimetro() float64
}

func Medidas(g Figuras) {
	fmt.Println("Medidas:", g)
	fmt.Println("Area:", g.area())
	fmt.Println("Perimetro:", g.perimetro())
}

package main

import (
	"paquetes/figuras"
)

func main() {
	/* fmt.Println()
	mensajes.Saludar()
	mensajes.Imprimir() */

	cua1 := figuras.Cuadrado{Lado: 8.5}
	figuras.Medidas(&cua1)

	cir1 := figuras.Circulo{Radio: 2.54}
	figuras.Medidas(&cir1)
}

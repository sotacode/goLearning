package mensajes

import "fmt"

func Saludar() {
	fmt.Println("Hola desde el paquete mensaje")
}

const Mensajes = "Message from const"

func Imprimir() {
	fmt.Println(Mensajes)
	privateFunc()
}

func privateFunc() {
	fmt.Println("from private func")
}

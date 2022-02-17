package main

import (
	"fmt"
	"net/http"
	"time"
)

func revisarServidor(servidor string) {
	_, error := http.Get(servidor)
	if error != nil {
		fmt.Println("No esta disponible el servidor")
	} else {
		fmt.Println(servidor, "esta disponible")
	}
}

func main() {

	servidores := []string{
		"http://udemy.com",
		"http://oregoom.com",
		"http://youtube.com",
		"http://facebook.com",
		"http://google.com",
	}
	inicio := time.Now()
	for _, servidor := range servidores {
		revisarServidor(servidor)
	}
	diffTime := time.Since(inicio)
	fmt.Println("Tiempo de ejecucion:", diffTime)
}

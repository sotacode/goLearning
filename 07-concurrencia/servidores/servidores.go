package main

import (
	"fmt"
	"net/http"
	"time"
)

func revisarServidor(servidor string, channel chan string) {
	_, error := http.Get(servidor)
	if error != nil {
		//fmt.Println("No esta disponible el servidor")
		channel <- servidor + "No esta disponible"
	} else {
		//fmt.Println(servidor, "esta disponible")
		channel <- servidor + "esta disponible"
	}
}

func main() {

	canal := make(chan string)

	servidores := []string{
		"http://udemy.com",
		"http://oregoom.com",
		"http://youtube.com",
		"http://facebook.com",
		"http://google.com",
	}
	inicio := time.Now()
	for _, servidor := range servidores {
		go revisarServidor(servidor, canal)
	}
	for i := 0; i < len(servidores); i++ {
		fmt.Println(<-canal)
	}
	diffTime := time.Since(inicio)
	fmt.Println("Tiempo de ejecucion:", diffTime)
}

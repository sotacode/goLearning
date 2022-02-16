package main

import "fmt"

func main() {

	//infinito
	/* for {
		fmt.Println("codigo ejecutado infinitamente")
	} */

	//while
	numeros := 1245
	c := 0
	for numeros > 0 {
		numeros /= 10
		c++
	}

	fmt.Println(c)
}

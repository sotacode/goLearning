package main

import "fmt"

func main() {
	//not
	fmt.Println(!true)

	//and
	fmt.Println(true && true)
	fmt.Println(true && false)

	//Or
	fmt.Println(true || false)
	fmt.Println(false || false)
}

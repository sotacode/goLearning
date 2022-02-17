package main

import "fmt"

type User struct {
	name  string
	email string
	state bool
}

type Student struct {
	user   User
	codigo string
}

func main() {

	sota := User{"Sota", "holarina1@gmail.com", false}

	isi := User{"isi", "isi@gmail.com", true}

	sotaStudent := Student{
		user:   sota,
		codigo: "0xsd54w",
	}
	fmt.Println(isi)
	fmt.Println(sotaStudent.user)
}

package main

import "fmt"

//relaciones one to one
/* type User struct {
	name  string
	email string
	state bool
}
type Student struct {
	user   User
	codigo string
} */

//relaciones one to much
type Curso struct {
	titulo string
	videos []Video
}

type Video struct {
	titulo string
	curso  Curso
}

func main() {

	/* sota := User{"Sota", "holarina1@gmail.com", false}

	isi := User{"isi", "isi@gmail.com", true}

	sotaStudent := Student{
		user:   sota,
		codigo: "0xsd54w",
	}
	fmt.Println(isi)
	fmt.Println(sotaStudent.user) */

	v1 := Video{titulo: "Intro"}
	v2 := Video{titulo: "Medium course"}
	v3 := Video{titulo: "finish course"}

	curso := Curso{
		titulo: "Curso de Go",
		videos: []Video{v1, v2, v3},
	}
	v1.curso = curso
	v2.curso = curso
	v3.curso = curso

	fmt.Println(v1.curso.titulo)
	for _, video := range curso.videos {
		fmt.Println(video.titulo)
	}
}

package main

import (
	"fmt"
	"sgr-go/models"
)

type x struct {
	nome string
	sexo string
}

func main() {

	p := models.Person{Name: "carlos", Profession: "marceneiro", Age: 22}

	fmt.Println(p.ToString())
}
